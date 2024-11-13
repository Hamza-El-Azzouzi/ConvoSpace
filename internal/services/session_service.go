package services

import (
	"time"

	"forum/internal/repositories"

	"github.com/gofrs/uuid/v5"
)

type SessionService struct {
	SessionRepo *repositories.SessionsRepositorie
}

func (s *SessionService) DeleteSession(sessionID string) error {
	err := s.SessionRepo.DeletSession(sessionID)
	return err
}

func (s *SessionService) CreateSession(sessionID string, expiration time.Time, userID uuid.UUID) error {
	err := s.SessionRepo.Createession(sessionID, expiration, userID)
	return err
}
func (s * SessionService) DeleteSessionByDate (time time.Time) error {
	err := s.SessionRepo.DeleteSessionByDate(time)
	return err
}
