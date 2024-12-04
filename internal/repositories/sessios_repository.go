package repositories

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid/v5"
)

type SessionsRepositorie struct {
	DB *sql.DB
}

// delete session by id used in lougout service
func (s *SessionsRepositorie) DeletSession(sessionID string) error {
	return nil
}

// creation a session for a user used in login service
func (s *SessionsRepositorie) Createession(sessionID string, expiration time.Time, userID uuid.UUID) error {
	return nil
}

// delete a session by date used in session service
// this function delete all expired session if expired_at < time.now()
func (s *SessionsRepositorie) DeleteSessionByDate(time time.Time) error {
	return nil
}

// get the user id by his session id
func (s *SessionsRepositorie) GetUser(sessionID string) (string, error) {
	return "", nil
}
