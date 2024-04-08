package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
)

func webauthnAuthenticationBegin(webauthn *webauthn.WebAuthn,
	sessionDataStore sessionDataStore,
) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		assertion, sd, err := webauthn.BeginDiscoverableLogin()
		if err != nil {
			handleError(ginContext, err, http.StatusInternalServerError)
			return
		}
		sessionDataStore.Add(*sd)
		ginContext.JSON(http.StatusOK, assertion.Response)
	}
}
