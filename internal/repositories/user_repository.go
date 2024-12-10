package repositories

import (
	"database/sql"

	"forum/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

// create a user in register form used in register service
func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (id, username, email, password_hash) VALUES (?,?,?,?)`
	_, err := r.DB.Exec(query, user.ID, user.Username, user.Email, user.PasswordHash)
	return err
}

func (repo *UserRepository) FindUser(identifier string, flag string) (*models.User, error) {
	user := &models.User{}
	query := ""
	switch true {
	case flag == "byId":
		query = `SELECT id, username, email, password_hash FROM users WHERE id= ?`
	case flag == "byEmail":
		query = `SELECT id, username, email, password_hash FROM users WHERE email= ?`
	}
	row := repo.DB.QueryRow(query, identifier)
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
	user := &models.User{}
	query := `SELECT users.id, users.username, users.email, users.password_hash
		FROM users 
		JOIN sessions ON users.id = sessions.user_id 
		WHERE sessions.session_id = ?`
	row := r.DB.QueryRow(query, sessionID)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
