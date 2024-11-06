package utils

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func OpenHtml(fileName string, w http.ResponseWriter, data any) {
	temp, err := template.ParseFiles("../../templates/" + fileName)
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

func SetupStaticFilesHandlers(response http.ResponseWriter, request *http.Request) {
	fileinfo, err := os.Stat("../../" + request.URL.Path)
	if !os.IsNotExist(err) && !fileinfo.IsDir() {
		http.FileServer(http.Dir("../../")).ServeHTTP(response, request)
	}
}
