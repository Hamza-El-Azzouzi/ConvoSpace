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

func (s *AuthService) Register(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &models.User{
		ID:           uuid.Must(uuid.NewV4()),
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	return s.UserRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil || user == nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		fmt.Printf("error f login : %v",err)
		return nil, err
	}

	return user, nil
}
