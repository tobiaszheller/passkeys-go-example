package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

func webauthnRegisterVerify(webauthn *webauthn.WebAuthn,
	userStore userStore,
	sessionDataStore sessionDataStore,
) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		parsedCCD, err := protocol.ParseCredentialCreationResponse(ginContext.Request)
		if err != nil {
			handleError(ginContext, err, http.StatusInternalServerError)
			return
		}
		challenge := parsedCCD.Response.CollectedClientData.Challenge
		sd := sessionDataStore.Get(challenge)
		if sd == nil {
			handleError(ginContext, errors.New("empty session data"), http.StatusInternalServerError)
			return
		}
		sessionDataStore.Delete(challenge)
		user := userStore.GetByID(string(sd.UserID))
		if user == nil {
			handleError(ginContext, errors.New("empty user"), http.StatusInternalServerError)
			return
		}
		cred, err := webauthn.CreateCredential(user, *sd, parsedCCD)
		if err != nil {
			handleError(ginContext, err, http.StatusInternalServerError)
			return
		}
		user.AddCredential(*cred)
		userStore.Upsert(user)
		ginContext.JSON(http.StatusOK, map[string]any{"verified": true})
	}
}
