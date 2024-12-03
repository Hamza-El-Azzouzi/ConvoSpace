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
	
}

func (a *AuthService) Login(email, password string) (*models.User, error) {
	
}

func (a *AuthService) CheckUserAlreadyLogged(userID uuid.UUID) ([]models.UserSession, error) {
	

func (a *AuthService) GetUserBySessionID(sessionID string) (*models.User, error) {
	
}