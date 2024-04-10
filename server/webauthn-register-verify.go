package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
)

func webauthnRegisterVerify(webauthn *webauthn.WebAuthn,
	userStore userStore,
	sessionDataStore sessionDataStore,
) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		// 1. parse from request.
		// 2. get sessionData from store.
		// 3. delete sessionData

		user := userStore.GetByID(string(sd.UserID))
		if user == nil {
			handleError(ginContext, errors.New("empty user"), http.StatusInternalServerError)
			return
		}

		// 4. create credetnail.
		// 5. Add credential to user
		userStore.Upsert(user)
		ginContext.JSON(http.StatusOK, map[string]any{"verified": true})
	}
}
