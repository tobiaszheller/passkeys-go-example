package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	webauthnPkg "github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
)

func webauthnAuthenticationVerify(webauthn *webauthnPkg.WebAuthn,
	userStore userStore,
	sessionDataStore sessionDataStore,
) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		parsedCRD, err := protocol.ParseCredentialRequestResponse(ginContext.Request)
		if err != nil {
			handleError(ginContext, err, http.StatusInternalServerError)
			return
		}
		challenge := parsedCRD.Response.CollectedClientData.Challenge
		sd := sessionDataStore.Get(challenge)
		if sd == nil {
			handleError(ginContext, errors.New("empty session data"), http.StatusInternalServerError)
			return
		}
		sessionDataStore.Delete(challenge)
		var user *User
		cred, err := webauthn.ValidateDiscoverableLogin(func(rawID, userHandle []byte) (webauthnPkg.User, error) {
			user = userStore.GetByID(string(userHandle))
			if user == nil {
				return user, fmt.Errorf("failed to get user")
			}
			return user, nil
		}, *sd, parsedCRD)
		if err != nil {
			handleError(ginContext, err, http.StatusInternalServerError)
			return
		}
		// update credentials instead adding one.
		user.AddCredential(*cred)
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
