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
	// ErrMessage string
}
type SignUpReply struct {
	REplyMssg string
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
			utils.Error(w, http.StatusBadRequest)
		}
		if !h.AuthMidlaware.IsValidEmail(info.Email) ||
			!h.AuthMidlaware.IsValidName(info.Username) ||
			!h.AuthMidlaware.IsValidPassword(info.Passwd) ||
			!h.AuthMidlaware.IsValidPassword(info.ConfirmPasswd) ||
			!h.AuthMidlaware.IsmatchPassword(info.Passwd, info.ConfirmPasswd) {
			utils.Error(w, http.StatusBadRequest)
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
			}
		}
		sendResponse(w, "Done")
	}
}
