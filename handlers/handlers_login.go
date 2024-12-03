package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type LoginData struct {
	Email  string `json:"email"`
	Passwd string `json:"password"`
	// ErrMessage string
}

type Reply struct {
	REplyMssg string
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	switch true {
	case r.URL.Path != "/login":
		HandleError(w, 404)
	case r.Method == http.MethodGet:
		tmpl, err := template.ParseFiles("templates/templates_login.html")
		if err != nil {
			HandleError(w, 404)
		}
		// fmt.Println("test ", info)
		err = tmpl.Execute(w, "info1")
		if err != nil {
			HandleError(w, 500)
		}
	case r.Method == http.MethodPost:
		var info LoginData
		err := json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			HandleError(w, 400)
		}
		fmt.Println("testDAta ", info)
		w.Header().Set("Content-Type", "application/json")
		response := &Reply{
			REplyMssg: "login Done",
		}
		json.NewEncoder(w).Encode(&response)
	}
}
