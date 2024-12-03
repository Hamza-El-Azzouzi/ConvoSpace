package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"Forum_izahid/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func ProcessDB() {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := os.ReadFile("data.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	// var db *sql.DB
	dataStat := strings.Split(string(data), ";")
	// fmt.Println("test", dataStat[0])

	for _, stat := range dataStat {
		_, err := db.Exec(stat)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("test ", )
	}
	_, err = db.Exec(`INSERT INTO users (username,email,passwd)  VALUES (?,?,?) `, "eman", "test.com", "test00")
	// _, err = db.Exec(`INSERT INTO users (username,email,passwd)  VALUES (?,?,?) `, "eman", "test.com", "test00")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	ProcessDB()
	http.HandleFunc("/", handlers.HandleSignUp)
	http.HandleFunc("/login", handlers.HandleLogin)
	http.HandleFunc("/index", handlers.HandleIndex)
	fmt.Println("http://localhost:6060/")
	http.ListenAndServe(":6060", nil)
}
