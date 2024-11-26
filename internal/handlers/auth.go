package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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
		utils.Error(w, http.StatusMethodNotAllowed)
	}
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	sessionID := cookie.Value

	err = h.SessionService.DeleteSession(sessionID)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError)
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

	errFrom := map[string]string{}

	if r.Method == http.MethodPost {
		var userData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		err := json.NewDecoder(r.Body).Decode(&userData)
		if err != nil {
			fmt.Println(err)
			utils.Error(w, http.StatusBadRequest)
			return
		}

		if userData.Email == "" {
			errFrom["email"] = "Email is required."
		} else {
			if !h.AuthMidlaware.IsValidEmail(userData.Email) {
				errFrom["email"] = "Invalid email!"
			}
		}

		if userData.Password == "" {
			errFrom["password"] = "Password is required."
		}

		if len(errFrom) > 0 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errFrom)
			return
		}
		user, loginError := h.AuthService.Login(userData.Email, userData.Password)

		if user == nil || loginError != nil {
			if strings.Contains(loginError.Error(), "email") {
				errFrom["email"] = "There is no account linked to this email!"
			} else {
				errFrom["password"] = "Wrong password!"
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errFrom)
			return
		}

		expiration := time.Now().Add(60 * time.Minute)
		sessionID := uuid.Must(uuid.NewV4()).String()
		err = h.SessionService.CreateSession(sessionID, expiration, user.ID)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			Expires:  expiration,
			HttpOnly: true,
		})
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("")
		return
	} else {
		utils.Error(w, http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandler) RegisterHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.OpenHtml("signup.html", w, nil)
		return
	}

	errFrom := map[string]string{}
	if r.Method == http.MethodPost {
		var userDataRegister struct {
			Username         string `json:"username"`
			Email            string `json:"email"`
			Password         string `json:"password"`
			ConfirmePassword string `json:"confirmPassword"`
		}
		err := json.NewDecoder(r.Body).Decode(&userDataRegister)
		if err != nil {
			fmt.Println(err)
			utils.Error(w, http.StatusBadRequest)
			return
		}

		if userDataRegister.Username == "" {
			errFrom["username"] = "Username is required."
		}
		if userDataRegister.Email == "" {
			errFrom["email"] = "Email is required."
		} else {
			if !h.AuthMidlaware.IsValidEmail(userDataRegister.Email) {
				errFrom["email"] = "Invalid email!"
			}
		}
		if userDataRegister.Password == "" {
			errFrom["password"] = "Password is required."
		} else {
			if !h.AuthMidlaware.IsValidPassword(userDataRegister.Password) {
				errFrom["password"] = "Invalid Password<br>At least 8 characters<br>Contains at least one letter<br>Contains at least one digit<br>Contains at least one special characte"
			}
		}
		if userDataRegister.ConfirmePassword == "" {
			errFrom["conPassowrd"] = "Confirm password is required."
		} else {
			if userDataRegister.Password != userDataRegister.ConfirmePassword {
				errFrom["conPassowrd"] = "Passwords do not match!"
			}
		}
		if len(errFrom) > 0 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errFrom)
			return

		}
		registrError := h.AuthService.Register(userDataRegister.Username, userDataRegister.Email, userDataRegister.Password)
		if registrError != nil {
			if registrError.Error() == "email already exist" {
				errFrom["email"] = "The Email Already Exist"
			} else {
				errFrom["username"] = "The Usernames Already Exist"
			}
		}
		if len(errFrom) > 0 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errFrom)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("")
	} else {
		utils.Error(w, http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandler) CheckDoubleLogging(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed)
	}

	isLogged, user := h.AuthMidlaware.IsUserLoggedIn(w, r)

	if isLogged {
		userSEssion, errSession := h.AuthService.UserRepo.CheckUserAlreadyLogged(user.ID)
		if errSession != nil {
			utils.Error(w, http.StatusInternalServerError)
		}

		if len(userSEssion) > 1 {
			http.Redirect(w, r, "/logout", http.StatusSeeOther)
			return
		}
	}
}
