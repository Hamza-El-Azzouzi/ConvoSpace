package repositories

import (
	"database/sql"

	"forum/internal/models"

	// "golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(user *models.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users (id, username, email, password_hash) VALUES (?, ?, ?, ?)",
		user.ID, user.Username, user.Email, user.PasswordHash,
	)
	return err
}

func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	query := `SELECT id, username, email, password_hash FROM users WHERE email = ?`
	user := &models.User{}

	row := repo.DB.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) FindByID(userId string) (*models.User, error) {
	query := `SELECT id, username, email, password_hash FROM users WHERE id = ?`
	user := &models.User{}

	row := repo.DB.QueryRow(query, userId)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// Implement other CRUD operations...
