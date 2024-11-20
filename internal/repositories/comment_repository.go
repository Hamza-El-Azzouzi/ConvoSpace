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

func (c *CommentRepositorie) GetCommentByPost(postID string) ([]models.CommentDetails, error) {
	var comments []models.CommentDetails
	rows, err := c.DB.Query(`SELECT 
		comments.id AS comment_id,
		comments.content,
		comments.created_at,
		users.id AS user_id,
		users.username,
		(SELECT COUNT(*) FROM likes WHERE likes.comment_id = comments.id AND likes.react_type = "like") AS likes_count,
		(SELECT COUNT(*) FROM likes WHERE likes.comment_id = comments.id AND likes.react_type = "dislike") AS dislike_count
		FROM 
			comments
		JOIN 
			users ON comments.user_id = users.id
		WHERE
			comments.post_id = ?
		GROUP BY 
			comments.id
		
		ORDER BY comments.created_at DESC;`, postID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var comment models.CommentDetails
		err = rows.Scan(
			&comment.CommentID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UserID,
			&comment.Username,
			&comment.LikeCountComment,
			&comment.DisLikeCountComment,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)

	}
	return comments, nil
}
