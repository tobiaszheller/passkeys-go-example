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
		// 1. Begin disc login
		// 2. add sessionData to store.

		ginContext.JSON(http.StatusOK, assertion.Response)
	}
}
