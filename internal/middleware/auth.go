package middleware

import (
	"net/http"
	"regexp"

	"forum/internal/models"
	"forum/internal/services"
)

type AuthMidlaware struct {
	AuthService *services.AuthService
}

func (h *AuthMidlaware) IsUserLoggedIn(w http.ResponseWriter, r *http.Request) (bool, *models.User) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return false, nil
	}

	sessionID := cookie.Value
	
	var userID string
	err = h.AuthService.UserRepo.DB.QueryRow("SELECT user_id FROM sessions WHERE session_id = ?", sessionID).Scan(&userID)
	if err != nil {
		return false, nil
	}

	user, err := h.AuthService.UserRepo.FindByID(userID)
	if err != nil || user == nil {
		return false, nil
	}

	return true, user
}

func (h *AuthMidlaware) IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func (h *AuthMidlaware) IsValidPassword(password string) bool {

	regex := `^([a-z0-9A-Z]).{7,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(password)
}
