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
// check username&email already exists if not hash the password and send all data  to userRepo
func (a *AuthService) Register(username, email, password string) error {
	
}
// get the user by email email compare the psawords
func (a *AuthService) Login(email, password string) (*models.User, error) {
	
}
// return all user session inside a slice by given a user ID
func (a *AuthService) CheckUserAlreadyLogged(userID uuid.UUID) ([]models.UserSession, error) {
	
// get a user data by a session ID given as paramtere
func (a *AuthService) GetUserBySessionID(sessionID string) (*models.User, error) {
	
}