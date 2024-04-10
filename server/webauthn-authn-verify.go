package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	webauthnPkg "github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
)

func webauthnAuthenticationVerify(webauthn *webauthnPkg.WebAuthn,
	userStore userStore,
	sessionDataStore sessionDataStore,
) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		// 1. Parse data from reuqest
		// 2. Get sessionData from datastore
		// 3. delete session data

		var user *User
		// 4. validate disc login
		// 5. add credential to user.
		userStore.Upsert(user)
		ginContext.JSON(http.StatusOK, map[string]any{
			"verified": true,
			"user":     userToUserInfo(user),
		})
	}
}

func userToUserInfo(in *User) userInfo {
	userCreds := make([]userCredentials, 0, len(in.credentials))
	for _, credential := range in.credentials {
		name := "unknown"
		aaguid, err := uuid.FromBytes(credential.Authenticator.AAGUID)
		if err == nil {
			if v, ok := aaguidMapping[aaguid.String()]; ok {
				name = v
			}
		}

		userCreds = append(userCreds, userCredentials{
			ID:   string(credential.ID),
			Name: name,
		})
	}
	return userInfo{
		Name:        in.name,
		Credentials: userCreds,
	}
}

type userInfo struct {
	Name        string            `json:"name,omitempty" :"name"`
	Credentials []userCredentials `json:"credentials,omitempty" :"credentials"`
}

type userCredentials struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
