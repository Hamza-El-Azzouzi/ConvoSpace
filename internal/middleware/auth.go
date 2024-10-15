package middleware

import (
	"net/http"

	"forum/internal/models"
	"forum/internal/services"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func (h *AuthHandler) IsUserLoggedIn(w http.ResponseWriter, r *http.Request) (bool, *models.User) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return false, nil
	}

	sessionID := cookie.Value

	// Query the database to find the session
	var userID string
	err = h.AuthService.UserRepo.DB.QueryRow("SELECT user_id FROM sessions WHERE session_id = ?", sessionID).Scan(&userID)
	if err != nil {
		return false, nil
	}

	// Retrieve user from database using userID
	user, err := h.AuthService.UserRepo.FindByID(userID)
	if err != nil || user == nil {
		return false, nil
	}

	return true, user
}
