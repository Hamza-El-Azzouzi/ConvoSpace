package repositories

import (
	"database/sql"

	"forum/internal/models"
)

type CommentRepositorie struct {
	DB *sql.DB
}

func (c *CommentRepositorie) Create(comment *models.Comment) error {
	_, err := c.DB.Exec(
		"INSERT INTO comments (id, user_id, post_id, Content) VALUES (?, ?, ?, ?)",
		comment.ID, comment.UserID, comment.PostID, comment.Content)
	return err
}
