package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"Forum_izahid/dataBase"
)

type SignUpData struct {
	Username string `json:"username"`
	Passwd   string `json:"password"`
	Email    string `json:"email"`
	// ErrMessage string
}
type SendReply struct {
	REplyMssg string
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/templates_index.html")
	if err != nil {
		HandleError(w, 404)
	}
	err = tmpl.Execute(w, "hello")
	if err != nil {
		HandleError(w, 500)
	}
}

func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	// test := "HELLO"
	switch true {
	case r.URL.Path != "/":
		HandleError(w, 404)
	case r.Method == http.MethodGet:
		// HandleError(w, 405)
		tmpl, err := template.ParseFiles("templates/templates_signUp.html")
		if err != nil {
			HandleError(w, 404)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			HandleError(w, 500)
		}
	case r.Method == http.MethodPost:
		var info SignUpData
		err := json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			HandleError(w, 400)
		}
		if !VerifyData(&info) {
			HandleError(w, 400)
		} else {
			fmt.Println("--------")
			_, err = dataBase.Db.Exec(`INSERT INTO users (username,email,passwd)  VALUES (?,?,?) `, info.Username, info.Email, info.Passwd)
			if err != nil {
				fmt.Println("--------error")
				HandleError(w, 500)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		response := &SendReply{
			REplyMssg: "Sign Done",
		}
		json.NewEncoder(w).Encode(&response)
	}
}
