package repositories

import (
	"database/sql"
	"fmt"

	"forum/internal/models"

	"github.com/gofrs/uuid/v5"
)

type LikeReposetorie struct {
	DB *sql.DB
}

func (l *LikeReposetorie) Create(like *models.Like, liked string) error {
	var existingID uuid.UUID
	var existingReactType string
	var reactOn *string
	if liked == "post" {
		reactOn = like.PostID
	} else {
		reactOn = like.CommentID
	}
	err := l.DB.QueryRow(
		"SELECT id, react_type FROM likes WHERE user_id = ? AND "+liked+"_id = ?",
		like.UserID, reactOn,
	).Scan(&existingID, &existingReactType)

	if err == nil {
		if existingReactType == like.ReactType {

			_, err = l.DB.Exec("DELETE FROM likes WHERE id = ?", existingID)
			if err != nil {
				return fmt.Errorf("failed to remove existing reaction: %w", err)
			}
			return nil
		} else {
			_, err = l.DB.Exec(
				"UPDATE likes SET react_type = ? WHERE id = ?",
				like.ReactType, existingID,
			)
			if err != nil {
				return fmt.Errorf("failed to update reaction type: %w", err)
			}
			return nil
		}
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("failed to check existing reaction: %w", err)
	}
	_, err = l.DB.Exec(
		"INSERT INTO likes (id, user_id, post_id, comment_id, react_type) VALUES (?, ?, ?, ?, ?)",
		like.ID, like.UserID, like.PostID, like.CommentID, like.ReactType,
	)
	if err != nil {
		return fmt.Errorf("failed to create new reaction: %w", err)
	}
	return nil
}
