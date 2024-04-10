package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
)

func webauthnRegisterBegin(webauthn *webauthn.WebAuthn,
	userStore userStore,
	sessionDataStore sessionDataStore,
) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		username, err := usernameFromRequest(ginContext.Request)
		if err != nil {
			handleError(ginContext, err, http.StatusBadRequest)
			return
		}
		user := userStore.GetByName(username)
		if user == nil {
			user = NewUser(username)
			userStore.Upsert(user)
		}
		// 1. being register
		// 2. add sessionData to store.

		ginContext.JSON(http.StatusOK, cc.Response)
	}
}

func usernameFromRequest(r *http.Request) (string, error) {
	type registerRequest struct {
		Username string `json:"username"`
	}
	var req registerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return "", err
	}
	if req.Username == "" {
		return "", errors.New("empty username")
	}
	return req.Username, nil
}
