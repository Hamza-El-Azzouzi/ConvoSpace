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

func (h *AuthHandler) LogoutHandle(w http.ResponseWriter, r *http.Request) {
}

func (h *AuthHandler) LoginHandle(w http.ResponseWriter, r *http.Request) {
}

func (h *AuthHandler) RegisterHandle(w http.ResponseWriter, r *http.Request) {
}

func (h *AuthHandler) CheckDoubleLogging(w http.ResponseWriter, r *http.Request) {
}
