package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"forum/internal/utils"

	"github.com/gofrs/uuid/v5"
)

type LoginData struct {
	Email  string `json:"email"`
	Passwd string `json:"password"`
	// ErrMessage string
}

type LoginReply struct {
	REplyMssg string
}

func (h *AuthHandler) LoginHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test00")
	ActiveUser, _ := h.AuthMidlaware.IsUserLoggedIn(w, r)
	fmt.Println("test01")
	if ActiveUser {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	switch true {
	case r.Method == http.MethodGet:
		utils.OpenHtml("templates_login.html", w, nil)
	case r.Method == http.MethodPost:
		var info LoginData
		// if ActiveUser {
		// 	// test = "in session"
		// 	sendResponse(w, "session")
		// 	return
		// }
		err := json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			utils.Error(w, http.StatusBadRequest)
			return
		}
		if !h.AuthMidlaware.IsValidEmail(info.Email) || !h.AuthMidlaware.IsValidPassword(info.Passwd) {
			utils.Error(w, http.StatusBadRequest)
			return
		} else {
			user, err := h.AuthService.Login(info.Email, info.Passwd)
			if err != nil || user == nil {
				switch true {
				case err.Error() == "in email":
					sendResponse(w, "email")
					return
				case err.Error() == "in password":
					sendResponse(w, "passdw")
					return
				}
				// utils.Error(w, http.StatusBadRequest)
			} else {
				sessionExpires := time.Now().Add(5 * 60 * time.Minute)
				sessionId := uuid.Must(uuid.NewV4()).String()
				fmt.Println("test02")
				userSession := h.SessionService.CreateSession(sessionId, sessionExpires, user.ID)
				fmt.Println("test03")
				if userSession != nil {
					utils.Error(w, http.StatusInternalServerError)
					return
				}
				SetCookies(w, "sessionId", sessionId, sessionExpires)
				sendResponse(w, "login Done")
			}
		}
	}
}

func sendResponse(w http.ResponseWriter, reply string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	response := &LoginReply{
		REplyMssg: reply,
	}
	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError)
		return
	}
}

func SetCookies(w http.ResponseWriter, name, value string, expires time.Time) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Expires:  expires,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
}

// expire the cookies time and delete sessionId from the session table
func (h *AuthHandler) LogoutHandle(w http.ResponseWriter, r *http.Request) {
	activeUser, _ := h.AuthMidlaware.IsUserLoggedIn(w, r)
	if !activeUser {
		utils.Error(w, http.StatusBadRequest)
		return
	}
	sessionId, err := r.Cookie("sessionId")
	if err == nil || sessionId.Value != "" {
		err := h.SessionService.DeleteSession(sessionId.Value)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
	}

	SetCookies(w, "sessionId", "", time.Now().Add(-1*time.Hour))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
