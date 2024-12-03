package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type SignUpData struct {
	Username string `json:"username"`
	Passwd   string `json:"password"`
	Email    string `json:"email"`
	// ErrMessage string
}
type Reply struct {
	REplyMssg string
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
			fmt.Println("test ", info)
		}
		w.Header().Set("Content-Type", "application/json")
		response := &Reply{
			REplyMssg: "Done",
		}
		json.NewEncoder(w).Encode(&response)
	}

	// fmt.Println("test ", info)
	// info := &SignUpData{
	// 	Username: r.FormValue("username"),
	// 	Email:    r.FormValue("email"),
	// 	Passwd:   r.FormValue("passwd"),
	// }

	// if !VerifyData(info) {
	// 	fmt.Println("errorData")
	// 	// info.ErrMessage = "invalid DATA"
	// 	// http.Redirect(w, r, "/", 200)
	// 	tmpl, err = template.ParseFiles("templates/templates_signUp.html")
	// 	if err != nil {
	// 		HandleError(w, 404)
	// 	}
	// }
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testLOGIN")
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
		tmpl, err := template.ParseFiles("templates/templates_login.html")
		if err != nil {
			HandleError(w, 404)
		}
		// fmt.Println("test ", info)
		err = tmpl.Execute(w, "info2")
		if err != nil {
			HandleError(w, 500)
		}

	}

	// http.Redirect(w, r, "/", 200)
}
