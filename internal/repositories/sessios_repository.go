package repositories

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid/v5"
)

type SessionsRepositorie struct {
	DB *sql.DB
}

func (s *SessionsRepositorie) DeletSession(sessionID string) error {
	
}

func (s *SessionsRepositorie) Createession(sessionID string, expiration time.Time, userID uuid.UUID) error {
	
}


func (s *SessionsRepositorie) DeleteSessionByDate(time time.Time) error {
	
}

func (s *SessionsRepositorie) GetUser(sessionID string) (string,error) {
	
}