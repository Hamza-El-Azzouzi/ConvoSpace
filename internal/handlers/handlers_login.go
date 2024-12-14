package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"forum/internal/models"
	"forum/internal/utils"

	"github.com/gofrs/uuid/v5"
)

type LoginData struct {
	Email  string `json:"email"`
	Passwd string `json:"password"`
}



func (h *AuthHandler) LoginHandle(w http.ResponseWriter, r *http.Request) {
	ActiveUser, _ := h.AuthMidlaware.IsUserLoggedIn(w, r)
	switch true {
	case r.Method == http.MethodGet:
		if ActiveUser {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		utils.OpenHtml("login.html", w, nil)
		return
	case r.Method == http.MethodPost:
		var info LoginData
		err := json.NewDecoder(r.Body).Decode(&info)
		defer r.Body.Close()
		if err != nil {
			utils.Error(w, http.StatusBadRequest)
			return
		}
		if !h.AuthMidlaware.IsValidEmail(info.Email) || !h.AuthMidlaware.IsValidPassword(info.Passwd) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user, err := h.AuthService.Login(info.Email, info.Passwd)
		if err != nil || user == nil {
			switch true {
			case err.Error() == "in email":
				sendResponse(w, "email")
				return
			case strings.Contains(err.Error(), "password"):
				sendResponse(w, "passdw")
				return
			}
		}
		sessionExpires := time.Now().Add(5 * 60 * time.Minute)
		sessionId := uuid.Must(uuid.NewV4()).String()
		userSession := h.SessionService.CreateSession(sessionId, sessionExpires, user.ID)
		if userSession != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		SetCookies(w, "sessionId", sessionId, sessionExpires)
		sendResponse(w, "Done")

	}
}

func sendResponse(w http.ResponseWriter, reply string) {
	w.Header().Set("Content-Type", "application/json")

	if reply != "Done" {
		w.WriteHeader(http.StatusBadRequest)
	}

	response := &models.LoginReply{
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

// LogoutHandle func expire the cookies time and delete sessionId from the session table
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
