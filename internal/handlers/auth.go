package handlers

import (
	"encoding/json"
	"fmt"
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

// expire the cookies time and delete sessionId from the session table
func (h *AuthHandler) LogoutHandle(w http.ResponseWriter, r *http.Request) {
}

type LoginData struct {
	Email  string `json:"email"`
	Passwd string `json:"password"`
	// ErrMessage string
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
type LoginReply struct {
	REplyMssg string
}

func (h *AuthHandler) LoginHandle(w http.ResponseWriter, r *http.Request) {
	switch true {
	case r.Method == http.MethodGet:
		utils.OpenHtml("templates_login.html", w, nil)
	case r.Method == http.MethodPost:
		var info LoginData
		err := json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			utils.Error(w, http.StatusBadRequest)
		}
		// fmt.Println("testDAta ", info)
		if !h.AuthMidlaware.IsValidEmail(info.Email) ||
			!h.AuthMidlaware.IsValidPassword(info.Passwd) {
			utils.Error(w, http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		response := &LoginReply{
			REplyMssg: "login Done",
		}
		json.NewEncoder(w).Encode(&response)
	}
}

func (h *AuthHandler) RegisterHandle(w http.ResponseWriter, r *http.Request) {
	switch true {
	case r.Method == http.MethodGet:
		utils.OpenHtml("templates_signUp.html", w, nil)
	case r.Method == http.MethodPost:
		var info SignUpData
		err := json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			utils.Error(w, http.StatusBadRequest)
		}
		fmt.Println("--------->", info)
		if !h.AuthMidlaware.IsValidEmail(info.Email) ||
			!h.AuthMidlaware.IsValidName(info.Username) ||
			!h.AuthMidlaware.IsValidPassword(info.Passwd) ||
			!h.AuthMidlaware.IsValidPassword(info.ConfirmPasswd) ||
			!h.AuthMidlaware.IsmatchPassword(info.Passwd, info.ConfirmPasswd) {

			utils.Error(w, http.StatusBadRequest)

			fmt.Println(h.AuthMidlaware.IsValidEmail(info.Email),
				h.AuthMidlaware.IsValidName(info.Username),
				h.AuthMidlaware.IsValidPassword(info.Passwd),
				h.AuthMidlaware.IsValidPassword(info.ConfirmPasswd),
				h.AuthMidlaware.IsmatchPassword(info.Passwd, info.ConfirmPasswd))
		} else {

			w.Header().Set("Content-Type", "application/json")
			response := &SignUpReply{
				REplyMssg: "Sign Done",
			}

			json.NewEncoder(w).Encode(&response)
		}

	}
}

// check the user has already a session
func (h *AuthHandler) CheckDoubleLogging(w http.ResponseWriter, r *http.Request) {
}
