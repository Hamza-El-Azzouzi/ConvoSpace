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
	_, err := s.DB.Exec("DELETE FROM sessions WHERE session_id = ?", sessionID)

	return err
}

func (s *SessionsRepositorie) Createession(sessionID string, expiration time.Time, userID uuid.UUID) error {
	_, err := s.DB.Exec("INSERT INTO sessions (session_id, user_id, expires_at) VALUES (?, ?, ?)", sessionID, userID, expiration)
	return err
}


func (s *SessionsRepositorie) DeleteSessionByDate(time time.Time) error {
	_, err := s.DB.Exec("DELETE FROM sessions WHERE expires_at < ?", time)
	return err
}