package handlers

import (
	"encoding/json"
	"net/http"

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
	if ActiveUser {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	switch true {
	case r.Method == http.MethodGet:
		utils.OpenHtml("templates_signUp.html", w, nil)
	case r.Method == http.MethodPost:
		var info SignUpData
		err := json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			utils.Error(w, http.StatusBadRequest)
		}
		if !h.AuthMidlaware.IsValidEmail(info.Email) ||
			!h.AuthMidlaware.IsValidName(info.Username) ||
			!h.AuthMidlaware.IsValidPassword(info.Passwd) ||
			!h.AuthMidlaware.IsValidPassword(info.ConfirmPasswd) ||
			!h.AuthMidlaware.IsmatchPassword(info.Passwd, info.ConfirmPasswd) {

			utils.Error(w, http.StatusBadRequest)
		} else {
			userExist := h.AuthService.Register(info.Username, info.Email, info.Passwd)
			if userExist != nil {
				sendResponse(w, "err")
			} else {
				sendResponse(w, "Sign Done")
			}
		}
	}
}
