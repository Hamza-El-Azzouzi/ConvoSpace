package utils

import "fmt"

import "net/http"


type ErrorStruct struct {
	Err      string
	Name     string
	Messages string
}

func Error(w http.ResponseWriter, status int) {
	var errorMap = map[int][]string{
		1: {"Oops! JavaScript not found", "JavaScript is disabled in your browser. Please enable it for the best experience."},
		400: {"Bad Request", "The server could not understand the request due to invalid syntax. Please check your input and try again."},
		403: {"Forbidden", "You do not have permission to access this resource."},
		404: {"Oops! Page not found", "The page you're looking for doesn't exist or has been moved."},
		405: {"Method Not Allowed", "The request method is not supported for the requested resource."},
		429: {"Too Many Requests", "You have sent too many requests in a given amount of time. Please try again later."},
		500: {"Something Went Wrong!", "Oops! It looks like our server encountered an unexpected issue. Don't worry, our team has been notified."},
	
	}
	var errorData ErrorStruct
	if(status == 1){
		w.WriteHeader(404)
	}else{
		w.WriteHeader(status)
	}
	
	for key, value := range errorMap {
		if key == status {
			if key ==1 {
				errorData.Err ="🙂"
			}else{
				errorData.Err = fmt.Sprint(key)
			}
			errorData.Name = value[0]
			errorData.Messages = value[1]
			OpenHtml("error.html", w, errorData)
		}
	}
	errorMap = nil
}
