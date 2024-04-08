package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
)

func webauthnRegisterBegin(webauthn *webauthn.WebAuthn,
	userStore userStore,
	sessionDataStore sessionDataStore,
) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		type registerRequest struct {
			Username string `json:"username"`
		}
		var req registerRequest
		err := json.NewDecoder(ginContext.Request.Body).Decode(&req)
		if err != nil {
			handleError(ginContext, err, http.StatusBadRequest)
			return
		}
		if req.Username == "" {
			ginContext.JSON(http.StatusBadRequest, map[string]string{"error": "missing username"})
			return
		}
		user := userStore.GetByName(req.Username)
		if user == nil {
			user = NewUser(req.Username)
			userStore.Upsert(user)
		}

		cc, sd, err := webauthn.BeginRegistration(user)
		if err != nil {
			handleError(ginContext, err, http.StatusInternalServerError)
			return
		}
		sessionDataStore.Add(*sd)
		ginContext.JSON(http.StatusOK, cc.Response)
	}
}
