package main

import (
	"log"
	"net/http"
)

func main() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}
	log.Println("Listing on 8080 http://localhost:8080")
	srv.ListenAndServe()
}
