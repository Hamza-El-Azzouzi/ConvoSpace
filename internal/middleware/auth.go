package middleware

import (
	"fmt"
	"net/http"
	"regexp"

	"forum/internal/models"
	"forum/internal/services"
	"forum/internal/utils"
)

type AuthMiddleware struct {
	AuthService *services.AuthService
}

func (h *AuthMiddleware) IsUserLoggedIn(w http.ResponseWriter, r *http.Request) (bool, *models.User) {
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

func (h *AuthMiddleware) IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func (h *AuthMiddleware) IsValidPassword(password string) bool {
	regex := `^([a-z0-9A-Z]).{7,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(password)
}

func (h *AuthMiddleware) CheckDoubleLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isLogged, user := h.IsUserLoggedIn(w, r)
		fmt.Printf("user logged in: %v\n", user)

		if !isLogged {
			fmt.Printf("no user logged in found\n")
			next.ServeHTTP(w, r) // Proceed to the next handler even if not logged in.
			return
		}

		userSessions, errSession := h.AuthService.CheckUserAlreadyLogged(user.ID)
		if errSession != nil {
			fmt.Printf("session error: %v\n", errSession)
			utils.Error(w, 500)
			return
		}

		fmt.Printf("user sessions: %v\n", userSessions)
		if len(userSessions) > 1 {
			cookie, _ := r.Cookie("session_id")
			sessionID := cookie.Value
			fmt.Printf("this session will be deleted: %v\n", sessionID)

			http.Redirect(w, r, "/logout", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
