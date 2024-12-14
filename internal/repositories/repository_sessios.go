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
	query := `DELETE FROM sessions WHERE session_id = ?`
	_, err := s.DB.Exec(query, sessionID)
	return err
}


func (s *SessionsRepositorie) Createession(sessionID string, expiration time.Time, userID uuid.UUID) error {
	query := `INSERT INTO sessions (session_id, user_id, expires_at) VALUES (?, ?, ?)`
	_, err := s.DB.Exec(query, sessionID, userID, expiration)
	return err
}


func (s *SessionsRepositorie) UpdateSession(sessionID string, expiration time.Time, userID uuid.UUID) error {
	query := `	UPDATE sessions SET session_id= ?, expires_at=? WHERE user_id= ?`
	_, err := s.DB.Exec(query, sessionID, expiration, userID)
	return err
}

func (s *SessionsRepositorie) DeleteSessionByDate(time time.Time) error {
	query := `DELETE FROM sessions WHERE expires_at < ?`
	_, err := s.DB.Exec(query, time)
	return err
}

func (s *SessionsRepositorie) GetUser(sessionID string) (string, error) {
	userId := ""
	query := `SELECT user_id FROM sessions WHERE session_id = ?`
	row := s.DB.QueryRow(query, sessionID)
	err := row.Scan(&userId)
	if err != nil {
		return "", err
	}
	if err == sql.ErrNoRows {
		return "", err
	}
	return userId, nil
}

func (s *SessionsRepositorie) CheckUserAlreadyLogged(UserID uuid.UUID) error {
	var sessionID string
	query := `SELECT session_id
	 FROM sessions
	WHERE  user_id = ?`
	row := s.DB.QueryRow(query, UserID)
	err := row.Scan(&sessionID)
	if err != nil {
		return err
	}
	return nil
}