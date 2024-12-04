package handlers

import (
	"net/http"

	"forum/internal/middleware"
	"forum/internal/services"
)

type AuthHandler struct {
	AuthService    *services.AuthService
	AuthMidlaware  *middleware.AuthMiddleware
	SessionService *services.SessionService
}
// expire the cookies time and delete sessionId from the session table
func (h *AuthHandler) LogoutHandle(w http.ResponseWriter, r *http.Request) {
}

// handle log-in
func (h *AuthHandler) LoginHandle(w http.ResponseWriter, r *http.Request) {
}

func (h *AuthHandler) RegisterHandle(w http.ResponseWriter, r *http.Request) {
}

// check the user has already a session
func (h *AuthHandler) CheckDoubleLogging(w http.ResponseWriter, r *http.Request) {
}
