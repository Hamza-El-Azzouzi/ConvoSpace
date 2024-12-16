package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"forum/internal/middleware"
	"forum/internal/services"
	"forum/internal/utils"
)

type AuthHandler struct {
	AuthService    *services.AuthService
	AuthMidlaware  *middleware.AuthMiddleware
	SessionService *services.SessionService
}

type SignUpData struct {
	Username      string `json:"username"`
	Passwd        string `json:"password"`
	Email         string `json:"email"`
	ConfirmPasswd string `json:"confirmPassword"`
}

func (h *AuthHandler) RegisterHandle(w http.ResponseWriter, r *http.Request) {
	ActiveUser, _ := h.AuthMidlaware.IsUserLoggedIn(w, r)
	switch true {
	case r.Method == http.MethodGet:
		if ActiveUser {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		utils.OpenHtml("signUp.html", w, nil)
		return
	case r.Method == http.MethodPost:
		if ActiveUser {
			sendResponse(w, "session")
			return
		}
		var info SignUpData
		err := json.NewDecoder(r.Body).Decode(&info)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		if !h.AuthMidlaware.IsValidEmail(info.Email) ||
			!h.AuthMidlaware.IsValidName(info.Username) ||
			!h.AuthMidlaware.IsValidPassword(info.Passwd) ||
			!h.AuthMidlaware.IsmatchPassword(info.Passwd, info.ConfirmPasswd) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		userExist := h.AuthService.Register(info.Username, info.Email, info.Passwd)
		if userExist != nil {
			switch true {
			case userExist.Error() == "email":
				sendResponse(w, "email")
				return
			case strings.Contains(userExist.Error(), "username"):
				sendResponse(w, "user")
				return
			default:
<<<<<<< HEAD
				sendResponse(w, "password")
			}
		}
		sendResponse(w, "Done")
=======
				sendResponse(w, "passwd")
			}
		}
		sendResponse(w, "Done")
		// add default case
>>>>>>> ae52f3b57173b0649e2f850a435f7266201bbee8
	default:
		utils.Error(w, http.StatusMethodNotAllowed)
	}
}
