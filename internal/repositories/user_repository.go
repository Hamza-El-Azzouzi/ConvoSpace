package repositories

import (
	"database/sql"
	"fmt"
	"html"

	"forum/internal/models"

	"github.com/gofrs/uuid/v5"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(user *models.User) error {
	var err error
	return err
}

func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	
}

func (repo *UserRepository) FindByID(userId string) (*models.User, error) {
	
}

func (r *UserRepository) GetUserBySessionID(sessionID string) (*models.User, error) {
	
}

func (r *UserRepository) CheckUserAlreadyLogged(UserID uuid.UUID) ([]models.UserSession, error) {

}
