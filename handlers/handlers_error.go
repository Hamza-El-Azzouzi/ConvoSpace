package handlers

import (
	"html/template"
	"net/http"
)

type errorType struct {
	StatusCode int
	Message    string
}

func HandleError(w http.ResponseWriter, code int) {
	var error errorType
	switch code {
	case 404:
		error.StatusCode = code
		error.Message = "Page Not Found"
	case 405:
		error.StatusCode = code
		error.Message = "Method not Allowd"
	case 500:
		error.StatusCode = code
		error.Message = "Internal Server Error"
	}
	w.WriteHeader(error.StatusCode)

	tmpl, err := template.ParseFiles("templates/templates_error.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, error)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
