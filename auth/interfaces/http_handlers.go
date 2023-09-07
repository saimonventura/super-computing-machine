package interfaces

import (
	"encoding/json"
	"net/http"
	"time"

	"super-computing-machine/auth/application"
	"super-computing-machine/auth/domain"
)

type AuthHandler struct {
	AuthService *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: service,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq domain.User

	// Decode the incoming login json
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.AuthService.Authenticate(loginReq.Email, loginReq.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Set the token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "access_token",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour),
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
