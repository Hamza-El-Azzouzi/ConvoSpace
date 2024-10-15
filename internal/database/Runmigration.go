package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
)

func RunMigrations(db *sql.DB) error {
	// Check if tables already exist
	tables := []string{"users", "posts", "comments", "categories", "post_categories", "likes"}
	for _, table := range tables {
		if tableExists(db, table) {
			return fmt.Errorf("table %s already exists", table)
		}
	}

	// If no tables exist, proceed with migrations
	migrationSQL, err := ioutil.ReadFile("internal/database/migrations/001_initial_schema.sql")
	if err != nil {
		return err
	}

	// Split the SQL file into individual statements
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