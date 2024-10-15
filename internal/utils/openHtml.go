package utils

import (
	"net/http"
	"os"
	"text/template"
)

func OpenHtml(fileName string, response http.ResponseWriter, data any) {
	temp, _ := template.ParseFiles("./templates/" + fileName)

	_ = temp.Execute(response, data)
}

func SetupStaticFilesHandlers(response http.ResponseWriter, request *http.Request) {
	fileinfo, err := os.Stat("." + request.URL.Path)
	if !os.IsNotExist(err) && !fileinfo.IsDir() {
		http.FileServer(http.Dir("./")).ServeHTTP(response, request)
	}
}
