package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func getPath() string {
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
	basePath := getPath()
	temp, err := template.ParseFiles(basePath + "templates/" + fileName)
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
	path := ""
	basePath := getPath()
	if basePath == "" {
		path = "/app/"
	} else {
		path = basePath
	}
	defer func() {
		if err := recover(); err != nil {
			Error(w, 404)
		}
	}()

	fileinfo, err := os.Stat(path + r.URL.Path)
	fmt.Println(fileinfo)
	fmt.Printf("err file serve : %v \n",err)
	fmt.Println(basePath)
	if !os.IsNotExist(err) && !fileinfo.IsDir() {
		http.FileServer(http.Dir(basePath)).ServeHTTP(w, r)
	} else {
		Error(w, 404)
	}
}
