package server

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	webauthnPkg "github.com/go-webauthn/webauthn/webauthn"
)

func Run() (err error) {
	webauthn, err := webauthnPkg.New(&webauthnPkg.Config{
		RPID:          "localhost",
		RPDisplayName: "localhost",
		RPOrigins:     []string{"http://localhost:8080"},
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			ResidentKey:      protocol.ResidentKeyRequirementRequired,
			UserVerification: protocol.VerificationRequired,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create new webauthn: %w", err)
	}

	inMemoryUserStore := NewInMemoryUserStore()
	registerSessionDataStore := NewInMemSessionDataStore()
	loginSessionDataStore := NewInMemSessionDataStore()

	router := gin.New()
	router.POST("/generate-registration-options", webauthnRegisterBegin(
		webauthn,
		inMemoryUserStore,
		registerSessionDataStore,
	))
	router.POST("/verify-registration", webauthnRegisterVerify(
		webauthn,
		inMemoryUserStore,
		registerSessionDataStore))
	router.GET("/generate-authentication-options", webauthnAuthenticationBegin(
		webauthn,
		loginSessionDataStore))
	router.POST("/verify-authentication", webauthnAuthenticationVerify(
		webauthn,
		inMemoryUserStore,
		loginSessionDataStore))

	router.GET("/health", func(context *gin.Context) {
		context.Status(http.StatusOK)
		return
	})
	router.StaticFS("/home", mustFS())
	return router.Run(":8080")
}

type userStore interface {
	GetByName(username string) *User
	GetByID(id string) *User
	Upsert(u *User)
}

type sessionDataStore interface {
	Get(challenge string) *webauthnPkg.SessionData
	Add(sd webauthnPkg.SessionData)
	Delete(challenge string)
}

func mustFS() http.FileSystem {
	sub, err := fs.Sub(assets, "public")

	if err != nil {
		panic(err)
	}

	return http.FS(sub)
}

func handleError(ginContext *gin.Context, err error, status int) {
	var protocolError *protocol.Error
	if errors.As(err, &protocolError) {
		log.Println(protocolError.DevInfo)
	}
	log.Println("Error during handling request: ", err)
	if status >= http.StatusInternalServerError {
		ginContext.JSON(status, map[string]string{"error": "Internal server error"})
		return
	}
	ginContext.JSON(status, map[string]string{"error": err.Error()})
}
