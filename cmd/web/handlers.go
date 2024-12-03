package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../../assets/templates/home.page.html")
	if err != nil {
		fmt.Println("test0")
		http.Error(w, err.Error(), 500)
		return
	}
	error := t.Execute(w, nil)
	if error != nil {
		fmt.Println("test1")
		http.Error(w, err.Error(), 500)
		return
	}
}
