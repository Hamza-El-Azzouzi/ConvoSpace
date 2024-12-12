package utils

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func GetPath() string {
	basePath := ""
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	if filepath.Dir(ex) == "/app" {
		basePath = ""
	} else {
		basePath = "../../"
	}
	return basePath
}

func OpenHtml(fileName string, w http.ResponseWriter, data any) {
	basePath := GetPath()
	temp, err := template.ParseFiles(basePath + "templates/template_" + fileName)
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	err = temp.Execute(w, data)
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
}

func SetupStaticFilesHandlers(w http.ResponseWriter, r *http.Request) {
	path := ""
	basePath := GetPath()
	if basePath == "" {
		path = "/app/"
	} else {
		path = basePath
	}
	defer func() {
		err := recover()
		if err != nil {
			Error(w, http.StatusNotFound)
		}
	}()

	fileinfo, err := os.Stat(path + r.URL.Path)
	if !os.IsNotExist(err) && !fileinfo.IsDir() {
		http.FileServer(http.Dir(basePath)).ServeHTTP(w, r)
	} else {
		Error(w, http.StatusNotFound)
	}
}
