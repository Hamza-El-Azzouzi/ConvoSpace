package handlers

import (
	"fmt"
	"net/http"
	"time"

	"forum/internal/services"
	"forum/internal/utils"

	"github.com/gofrs/uuid/v5"
)

type AuthHandler struct {
	AuthService *services.AuthService
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
		utils.Error(w,500)
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
		utils.OpenHtml("login.html", w, r)
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			utils.Error(w,500)
		}

		email := r.Form.Get("email")
		password := r.Form.Get("password")

		if email == "" || password == "" {
			utils.Error(w,400)
		}

		user, loginError := h.AuthService.Login(email, password)
		if user == nil {
			fmt.Fprintln(w, "Invalid email or password")
			return
		}
		sessionID := uuid.Must(uuid.NewV4()).String()
		_, err = h.AuthService.UserRepo.DB.Exec("INSERT INTO sessions (session_id, user_id) VALUES (?, ?)", sessionID, user.ID)
		if err != nil {
			utils.Error(w,500)
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})

		fmt.Println(user)
		if loginError != nil {
			utils.Error(w,500)
		}
		fmt.Println("User loged successfully!")

		http.Redirect(w, r, "/", http.StatusSeeOther) 
	} else {
		utils.OpenHtml("login.html", w, r)
	}
}

func (h *AuthHandler) RegisterHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.OpenHtml("signup.html", w, r)
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			utils.Error(w,500)
		}

		userName := r.Form.Get("Username")
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		confirmPassword := r.Form.Get("Confirm-Password")

		if password != confirmPassword {
			fmt.Fprintln(w, "Passwords do not match!")
			return
		}

		if userName == "" || email == "" || password == "" {
			utils.Error(w,405)
		}

		registrError := h.AuthService.Register(userName, email, password)
		if registrError != nil {
			utils.Error(w,500)
		}
		fmt.Println("User registered successfully!")

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}