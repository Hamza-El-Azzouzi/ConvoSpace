package services

import (
	"fmt"

	"forum/internal/models"
	"forum/internal/repositories"

	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func (a *AuthService) Register(username, email, password string) error {
	user, err := a.UserRepo.FindByEmail(email)
	if user != nil || err != nil {
		return fmt.Errorf("email already exist")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user = &models.User{
		ID:           uuid.Must(uuid.NewV4()),
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	return a.UserRepo.Create(user)
}

func (a *AuthService) Login(email, password string) (*models.User, error) {
	user, err := a.UserRepo.FindByEmail(email)
	if err != nil || user == nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *AuthService) CheckUserAlreadyLogged(userID uuid.UUID) ([]models.UserSession, error) {
	return a.UserRepo.CheckUserAlreadyLogged(userID)
}

func (a *AuthService) GetUserBySessionID(sessionID string) (*models.User, error) {
	return a.UserRepo.GetUserBySessionID(sessionID)
}