package utils

import (
	"fmt"
	"log"
	"time"

	"forum/internal/services"
)

type Cleaner struct {
	SessionService *services.SessionService
}

func (c *Cleaner) CleanupExpiredSessions() {
	for {
		time.Sleep(1 * time.Minute)
		fmt.Println("working...")
		time := time.Now()
		err := c.SessionService.DeleteSessionByDate(time)
		if err != nil {
			log.Printf("Error deleting expired sessions: %v", err)
		}
	}
}
