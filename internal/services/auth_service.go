package services

import (
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
    // Retrieve the user by email
    user, err := s.UserRepo.FindByEmail(email)
    if err != nil {
        return nil, err // Return error if user is not found
    }

    // Compare the provided password with the stored hashed password
    err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
    if err != nil {
        return nil, err // Return error if password does not match
    }

    return user, nil // Return the user if authentication is successful
}



// Implement other auth-related functions...
