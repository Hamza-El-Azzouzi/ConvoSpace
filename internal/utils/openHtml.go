package utils

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

func OpenHtml(fileName string, response http.ResponseWriter, data any) {
	temp, err := template.ParseFiles("../../templates/" + fileName)
	if err != nil {
		log.Printf("error parsing template file: %v", err)
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := temp.Execute(response, data); err != nil {
		log.Printf("error executing template: %v", err)
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
	}
}

func SetupStaticFilesHandlers(response http.ResponseWriter, request *http.Request) {
	fileinfo, err := os.Stat("../../" + request.URL.Path)
	if !os.IsNotExist(err) && !fileinfo.IsDir() {
		http.FileServer(http.Dir("../../")).ServeHTTP(response, request)
	}
}
