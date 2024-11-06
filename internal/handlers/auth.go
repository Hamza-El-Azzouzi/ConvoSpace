package handlers

import (
	"fmt"
	"net/http"
	"time"

	"forum/internal/middleware"
	"forum/internal/services"
	"forum/internal/utils"

	"github.com/gofrs/uuid/v5"
)

type AuthHandler struct {
	AuthService   *services.AuthService
	AuthMidlaware *middleware.AuthMidlaware
}

func (h *AuthHandler) LogoutHandle(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	sessionID := cookie.Value

	_, err = h.AuthService.UserRepo.DB.Exec("DELETE FROM sessions WHERE session_id = ?", sessionID)
	if err != nil {
		utils.Error(w, 500)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *AuthHandler) LoginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.OpenHtml("login.html", w, nil)
		return
	}
	errFrom := make(map[string]string)
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			utils.Error(w, 500)
		}

		email := r.Form.Get("email")
		password := r.Form.Get("password")

		if email == "" || password == "" {
			errFrom["password"] = "They can't be Empty"
			utils.OpenHtml("login.html", w, errFrom)
			return
		}

		if !h.AuthMidlaware.IsValidEmail(email) {
			errFrom["email"] = "Invalid email"
		}
		if len(errFrom) == 1 {
			utils.OpenHtml("login.html", w, errFrom)
			return
		}
		user, loginError := h.AuthService.Login(email, password)
		if user == nil {
			errFrom["password"] = "Invalid email or password"
			utils.OpenHtml("login.html", w, errFrom)
			return
		}
		sessionID := uuid.Must(uuid.NewV4()).String()
		_, err = h.AuthService.UserRepo.DB.Exec("INSERT INTO sessions (session_id, user_id) VALUES (?, ?)", sessionID, user.ID)
		if err != nil {
			utils.Error(w, 500)
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})
		if loginError != nil {
			utils.Error(w, 500)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		utils.Error(w, http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandler) RegisterHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.OpenHtml("signup.html", w, nil)
		return
	}
	errFrom := make(map[string]string)
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			utils.Error(w, 500)
		}

		userName := r.Form.Get("Username")
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		confirmPassword := r.Form.Get("Confirm-Password")

		if password != confirmPassword {
			errFrom["conPassowrd"] = "Passwords do not match!"
		}
		if !h.AuthMidlaware.IsValidEmail(email) {
			errFrom["email"] = "Invalid email"
		}
		if !h.AuthMidlaware.IsValidPassword(password) {
			errFrom["password"] = "Invalid Password, At least 8 charachters"
		}
		if userName == "" || email == "" || password == "" {
			errFrom["empty"] = "The Fields can't be Empty"
		}
		if len(errFrom) > 0 {
			utils.OpenHtml("signup.html", w, errFrom)
			return
		}
		registrError := h.AuthService.Register(userName, email, password)
		if registrError != nil {
			fmt.Println(registrError)
			errFrom["alreadyExist"] = "The Email Already Exist"
			utils.OpenHtml("signup.html", w, errFrom)
			
			return
		}
	}
}
