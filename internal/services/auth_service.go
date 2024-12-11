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

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

func HashPassword(psswd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(psswd), bcrypt.DefaultCost)
	return string(bytes), err
}

// check username&email already exists if not hash the password and send all data  to userRepo
func (a *AuthService) Register(username, email, password string) error {
	checkByEmail, err := a.UserRepo.FindUser(email, "byEmail")
	// specify the error if username already exists or the email
	if checkByEmail != nil || err != nil {
		return fmt.Errorf("user already exist")
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

// get the user by email email compare the psawords
func (a *AuthService) Login(email, password string) (*models.User, error) {
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

// get a user data by a session ID given as paramtere
func (a *AuthService) GetUserBySessionID(sessionID string) (*models.User, error) {
	user, err := a.UserRepo.GetUserBySessionID(sessionID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
