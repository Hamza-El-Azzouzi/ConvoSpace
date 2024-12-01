package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

type SignUpData struct {
	Username   string
	Passwd     string
	Email      string
	ErrMessage string
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(("templates/templates_index.html"))
	if err != nil {
		HandleError(w, 404)
	}
	err = tmp.Execute(w, nil)
	if err != nil {
		HandleError(w, 500)
	}
}

func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	// test := "HELLO"
	switch true {
	case r.Method != http.MethodGet:
		HandleError(w, 405)
	case r.URL.Path != "/register":
		HandleError(w, 404)
	}
	tmpl, err := template.ParseFiles("templates/templates_signUp.html")
	if err != nil {
		HandleError(w, 404)
	}
	// fmt.Println("test ", info)
	err = tmpl.Execute(w, nil)
	if err != nil {
		HandleError(w, 500)
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	switch true {
	case r.Method != http.MethodPost:
		HandleError(w, 405)
	case r.URL.Path != "/login":
		HandleError(w, 404)
	}
	// tmpl, err := template.ParseFiles("templates/templates_login.html")
	// if err != nil {
	// 	HandleError(w, 404)
	// }
	info := &SignUpData{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Passwd:   r.FormValue("passwd"),
	}

	if !VerifyData(info) {
		fmt.Println("errorData")
		info.ErrMessage = "invalid DATA"
		// http.Redirect(w, r, "/", 200)
		// tmpl, err = template.ParseFiles("templates/templates_signUp.html")
		// if err != nil {
		// 	HandleError(w, 404)
		// }
	}
	// fmt.Println("test ", info)
	// err = tmpl.Execute(w, info)
	// if err != nil {
	// 	HandleError(w, 500)
	// }
	http.Redirect(w, r, "/", 200)
}
