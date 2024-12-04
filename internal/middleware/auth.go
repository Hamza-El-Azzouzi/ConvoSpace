package middleware

import (
	"net/http"
	"regexp"

	"forum/internal/models"
	"forum/internal/services"
)

type AuthMiddleware struct {
	AuthService    *services.AuthService
	SessionService *services.SessionService
}

const (
	ExpEmail = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// ExpPasswd = `^[A-Za-z0-9]{8,20}$`
	ExpName = `^[a-zA-Z0-9_]{3,20}$`
)

// check if the user is logged retunr true/false & userData || nil
func (h *AuthMiddleware) IsUserLoggedIn(w http.ResponseWriter, r *http.Request) (bool, *models.User) {
	return false, nil
}

func (h *AuthMiddleware) IsValidEmail(email string) bool {
	isValid, _ := regexp.MatchString(ExpEmail, email)
	return isValid
}

func (h *AuthMiddleware) IsValidName(username string) bool {
	isValid, _ := regexp.MatchString(ExpName, username)
	return isValid
}

func (h *AuthMiddleware) IsmatchPassword(password string, confirmPassword string) bool {
	match := password == confirmPassword
	return match
}

func (h *AuthMiddleware) IsValidPassword(password string) bool {
	secure := true

	ExpPasswd := []string{".{7,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
	for _, test := range ExpPasswd {
		isValid, _ := regexp.MatchString(test, password)
		if !isValid {
			secure = false
			break
		}
	}

	return secure
}

// midlware to check if the user has already an session opend
func (h *AuthMiddleware) CheckDoubleLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
