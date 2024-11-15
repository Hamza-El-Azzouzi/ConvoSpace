package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/gofrs/uuid/v5"
)

func RunMigrations(db *sql.DB) error {
	allExist := true
	tables := []string{"users", "posts", "comments", "categories", "post_categories", "likes"}
	for _, table := range tables {
		if !tableExists(db, table) {
			allExist = false
			break
		}
	}

	if allExist {
		return nil
	}

	migrationSQL, err := os.ReadFile("internal/database/migrations/001_initial_schema.sql")
	if err != nil {
		return err
	}

	statements := strings.Split(string(migrationSQL), ";")

	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		_, err = db.Exec(stmt)
		if err != nil {
			return err
		}
	}

	return nil
}

func tableExists(db *sql.DB, tableName string) bool {

	query := `SELECT name FROM sqlite_master WHERE type='table' AND name=?;`
	var name string
	err := db.QueryRow(query, tableName).Scan(&name)
	return err == nil
}

func InsertDefaultCategories(db *sql.DB) error {
	
	categories := []string{
		"Announcements",
		"General Discussion",
		"Q&A and Help",
		"Web Development",
		"Mobile App Development",
		"Game Development",
		"Cybersecurity",
		"Technology & Gadgets",
		"Coding Challenges & Projects",
		"Design and UI/UX",
		"Career Advice and Freelancing",
		"Feedback and Suggestions",
		"Off-Topic Lounge",
		"Events and Meetups",
	}

	for _, category := range categories {
		ID := uuid.Must(uuid.NewV4())

		var exists int
		err := db.QueryRow("SELECT COUNT(*) FROM categories WHERE name = ?", category).Scan(&exists)
		if err != nil {
			return fmt.Errorf("error checking category %s: %v", category, err)
		}

		if exists == 0 {
			_, err := db.Exec("INSERT INTO categories (ID , name) VALUES (?,?)", ID, category)
			if err != nil {
				return fmt.Errorf("error inserting category %s: %v", category, err)
			}
		}
	}
	return nil
}
