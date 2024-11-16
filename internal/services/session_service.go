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
	return s.SessionRepo.DeletSession(sessionID)
}

func (s *SessionService) CreateSession(sessionID string, expiration time.Time, userID uuid.UUID) error {
	return s.SessionRepo.Createession(sessionID, expiration, userID)
}
func (s * SessionService) DeleteSessionByDate(time time.Time) error {
	return s.SessionRepo.DeleteSessionByDate(time)
}

func (s *SessionService) GetUserService(sessionId string)(string,error){
	userID, err := s.SessionRepo.GetUser(sessionId)
	if err != nil{
		return "" , err
	}
	return userID, nil
	
}