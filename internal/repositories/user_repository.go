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
func (r *UserRepository) GetUserBySessionID(sessionID string) (*models.User, error) {
	// Prepare a new User instance
	user := &models.User{}

	// Query to fetch the username based on session ID
	err := r.DB.QueryRow(`
		SELECT users.id, users.username, users.email, users.password_hash
		FROM users 
		JOIN sessions ON users.id = sessions.user_id 
		WHERE sessions.session_id = ?`, sessionID).
		Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	
	// Check for errors and handle no rows found
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err // Other errors
	}

	return user, nil // Successful fetch
}
// Implement other CRUD operations...
