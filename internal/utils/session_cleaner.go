package utils

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)
type Cleaner struct{
	Db *sql.DB
}

func (c *Cleaner)CleanupExpiredSessions() {
	for {
		time.Sleep(1 * time.Minute)
		fmt.Println("i'm working now...")
		_, err := c.Db.Exec("DELETE FROM sessions WHERE expires_at < ?",time.Now())
		if err != nil {
			log.Printf("Error deleting expired sessions: %v", err)
		}
	}
}
