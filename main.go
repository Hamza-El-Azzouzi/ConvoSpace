package main

import (
	"fmt"
	"net/http"

	"Forum_izahid/dataBase"
	"Forum_izahid/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dataBase.ProcessDB()
	http.HandleFunc("/", handlers.HandleSignUp)
	http.HandleFunc("/login", handlers.HandleLogin)
	http.HandleFunc("/index", handlers.HandleIndex)
	fmt.Println("http://localhost:6060/")
	http.ListenAndServe(":6060", nil)
}
