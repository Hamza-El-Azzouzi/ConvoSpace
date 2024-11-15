package repositories

import (
	"database/sql"
	"fmt"

	"forum/internal/models"

	"github.com/gofrs/uuid/v5"
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
	user := &models.User{}

	err := r.DB.QueryRow(`
		SELECT users.id, users.username, users.email, users.password_hash
		FROM users 
		JOIN sessions ON users.id = sessions.user_id 
		WHERE sessions.session_id = ?`, sessionID).
		Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) CheckUserAlreadyLogged(UserID uuid.UUID) ([]models.UserSession, error) {
	var userSessions []models.UserSession
	query := `SELECT session_id ,user_id FROM sessions WHERE user_id = ?`
	rows, err := r.DB.Query(query, UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	for rows.Next() {
		var userSession models.UserSession
		err = rows.Scan(&userSession.ID, &userSession.USerID);
		if err != nil {
			return nil, fmt.Errorf("error scanning sessions with user info filter: %v", err)
		}
		userSessions = append(userSessions, userSession)
	}
	return userSessions, nil
}
