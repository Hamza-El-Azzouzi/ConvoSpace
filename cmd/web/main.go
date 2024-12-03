package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	exécution()
	mux := http.NewServeMux()
	mux.HandleFunc("/", getHome)
	log.Println("Listing on 8080 http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func exécution() {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	sqlUser := `CREATE TABLE IF NOT EXISTS
    users (
        id TEXT PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT (DATETIME ('now', 'localtime'))
    );`
	sqlposts := `
	CREATE TABLE IF NOT EXISTS
    posts (
        id TEXT PRIMARY KEY,
        user_id TEXT NOT NULL,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT (DATETIME ('now', 'localtime')),
        FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
    );`
	sqlComments := `
	CREATE TABLE IF NOT EXISTS
    comments (
        id TEXT PRIMARY KEY,
        user_id TEXT NOT NULL,
        post_id TEXT NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT (DATETIME ('now', 'localtime')),
        FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
        FOREIGN KEY (post_id) REFERENCES posts (id) ON UPDATE CASCADE ON DELETE CASCADE
    );`
	sqlLikes := `
	CREATE TABLE IF NOT EXISTS
    likes (
        id TEXT PRIMARY KEY,
        user_id TEXT NOT NULL,
        post_id TEXT,
        comment_id TEXT,
        react_type TEXT,
        created_at TIMESTAMP DEFAULT (DATETIME ('now', 'localtime')),
        FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
        FOREIGN KEY (post_id) REFERENCES posts (id) ON UPDATE CASCADE ON DELETE CASCADE,
        FOREIGN KEY (comment_id) REFERENCES comments (id) ON UPDATE CASCADE ON DELETE CASCADE,
        CONSTRAINT unique_user_post_comment UNIQUE (user_id, post_id, comment_id),
        CHECK (
            (
                post_id IS NOT NULL
                AND comment_id IS NULL
            )
            OR (
                post_id IS NULL
                AND comment_id IS NOT NULL
            )
        )
    );`
	var sl []string
	sl = append(sl, sqlUser, sqlposts, sqlComments, sqlLikes)
	for j := 0; j < len(sl); j++ {
		_, err = db.Exec(sl[j])
	}
	if err != nil {
		log.Fatalln(err)
	}
}
