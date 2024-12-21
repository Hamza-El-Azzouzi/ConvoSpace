package services

import (
	"fmt"
	"strings"

	"forum/internal/models"
	"forum/internal/repositories"

	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func HashPassword(psswd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(psswd), bcrypt.DefaultCost)
	return string(bytes), err
}

func (a *AuthService) Register(username, email, password string) error {
	email = strings.ToLower(email)
	username = strings.ToLower(username)
	checkByEmail, err := a.UserRepo.FindUser(email, "byEmail")
	if checkByEmail != nil {
		return fmt.Errorf("email")
	}
	if err != nil {
		return fmt.Errorf("user")
	}

	hash, err := HashPassword(password)
	if err != nil {
		return err
	}
	uuid := uuid.Must(uuid.NewV4())
	user := &models.User{
		ID:           uuid,
		Username:     username,
		Email:        email,
		PasswordHash: hash,
	}
	err = a.UserRepo.Create(user)
	return err
}

func (a *AuthService) Login(email, password string) (*models.User, error) {
	email = strings.ToLower(email)
	userByEmail, err := a.UserRepo.FindUser(email, "byEmail")
	if userByEmail == nil || err != nil {
		return nil, fmt.Errorf("in email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userByEmail.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return userByEmail, nil
}

func (a *AuthService) GetUserBySessionID(sessionID string) (*models.User, error) {
	user, err := a.UserRepo.GetUserBySessionID(sessionID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
