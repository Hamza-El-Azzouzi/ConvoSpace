package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

type SignUpData struct {
	Username string
	Passwd   string
	Email    string
}

func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	test := "HELLO "
	if r.Method != http.MethodPost {
		fmt.Println("")
		return
	}
	tmpl, err := template.ParseFiles("template_signUp.html")
	info := SignUpData{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
	}
	if err != nil {
		HandleError(w, 404)
	}
	err = tmpl.Execute(w, test)
	if err != nil {
		HandleError(w, 500)
	}
}
