package utils

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func OpenHtml(fileName string, w http.ResponseWriter, data any) {
	temp, err := template.ParseFiles("templates/" + fileName)
	if err != nil {
		fmt.Printf("error parsing template file: %v", err)
		Error(w, 500)
		return
	}
	err = temp.Execute(w, data)
	if err != nil {
		fmt.Printf("error executing template: %v", err)
		Error(w, 500)
		return
	}
}

func SetupStaticFilesHandlers(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if err := recover(); err != nil {
            Error(w, 500)
        }
    }()
    fileinfo, err := os.Stat("static" + r.URL.Path)
    if !os.IsNotExist(err) && !fileinfo.IsDir() {
		http.StripPrefix("/static/", http.FileServer(http.Dir("/app/static")))
    } else {
        Error(w, 404)
    }
}
