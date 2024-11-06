package utils

import "net/http"


type ErrorStruct struct {
	Err      int
	Name     string
	Messages string
}

func Error(response http.ResponseWriter, status int) {
	var errorMap = map[int][]string{
		400: {"Bad Request", "The server could not understand the request due to invalid syntax. Please check your input and try again."},
		404: {"Oops! Page not found", "The page you're looking for doesn't exist or has been moved."},
		405: {"Method Not Allowed", "The request method is not supported for the requested resource."},
		500: {"Something Went Wrong!", "Oops! It looks like our server encountered an unexpected issue. Don't worry, our team has been notified."},
		429: {"Too Many Requests", "You have sent too many requests in a given amount of time. Please try again later."},
	}
	var errorData ErrorStruct
	// response.WriteHeader(status)
	for key, value := range errorMap {
		if key == status {
			errorData.Err = key
			errorData.Name = value[0]
			errorData.Messages = value[1]
			OpenHtml("error.html", response, errorData)
		}
	}
	errorMap = nil
}
