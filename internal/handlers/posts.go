package handlers

import (
	"net/http"

	"forum/internal/utils"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"LoggedIn": true,
	}
	cookie, err := r.Cookie("session_id")
	if err != nil {
		data = map[string]interface{}{
			"LoggedIn": false,
		}
	}

	if cookie != nil {
		data["LoggedIn"] = true
	}
	// Retrieve user from database using userID
	utils.OpenHtml("index.html", w, data)
}
