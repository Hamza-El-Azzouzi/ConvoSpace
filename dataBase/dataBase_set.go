package dataBase

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

var Db *sql.DB

func ProcessDB() {
	var err error
	Db, err = sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer Db.Close()
	err = Db.Ping()
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
		if stat = strings.TrimSpace(stat); stat == "" {
			continue
		}

		_, err := Db.Exec(stat)
		if err != nil {
			fmt.Println("Error executing SQL statement:", err)
		}
		// fmt.Println("test ", )
	}
	// _, err = db.Exec(`INSERT INTO users (username,email,passwd)  VALUES (?,?,?) `, "eman", "test.com", "test00")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
