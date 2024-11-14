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
	AuthService    *services.AuthService
	AuthMidlaware  *middleware.AuthMiddleware
	SessionService *services.SessionService
}

func (h *AuthHandler) LogoutHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, 405)
	}
	cookie, _ := r.Cookie("session_id")

	sessionID := cookie.Value
	fmt.Println(sessionID)
	err := h.SessionService.DeleteSession(sessionID)
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

	http.Redirect(w, r, "/login", http.StatusSeeOther)
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
			utils.OpenHtml("login.html", w, errFrom)
			return
		}
		user, loginError := h.AuthService.Login(email, password)
		if loginError != nil {
			errFrom["password"] = "Invalid email or password"
		}
		if user == nil {
			errFrom["password"] = "Invalid email or password"
			utils.OpenHtml("login.html", w, errFrom)
			return
		}
		expiration := time.Now().Add(60 * time.Minute)
		sessionID := uuid.Must(uuid.NewV4()).String()
		err = h.SessionService.CreateSession(sessionID, expiration, user.ID)
		if err != nil {
			fmt.Printf("errr f creation sesions : %v \n", err)
			utils.Error(w, 500)
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			Expires:  expiration,
			HttpOnly: true,
		})
		fmt.Printf("user is logged with this seession : %v\n", sessionID)
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
		h.LoginHandle(w, r)
	}
}

func (h *AuthHandler) CheckDoubleLogging(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Error(w, 405)
	}

	isLogged, user := h.AuthMidlaware.IsUserLoggedIn(w, r)
	if !isLogged {
		fmt.Println("no user found logged")
	} else {
		userSEssion, errSession := h.AuthService.UserRepo.CheckUserAlreadyLogged(user.ID)
		if errSession != nil {
			fmt.Printf("err f session login : %v", errSession)
		}
		cookie, _ := r.Cookie("session_id")

		sessionID := cookie.Value
		
		if len(userSEssion) > 1 {
			fmt.Printf("this sessions will deleted : %v\n",sessionID)
			h.LogoutHandle(w, r)
			return
		}
	}
}
