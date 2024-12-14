package repositories

import (
	"database/sql"
	"html"

	"forum/internal/models"
)

type CommentRepositorie struct {
	DB *sql.DB
}

// save the comment in the data base and return error
func (c *CommentRepositorie) Create(comment *models.Comment) error {
	query := "INSERT INTO comments (id, user_id, post_id, content, created_at) VALUES (?, ?, ?, ?, ?)"
	prp, prepareErr := c.DB.Prepare(query)
	if prepareErr != nil {
		return prepareErr
	}
	defer prp.Close()
	comment.Content = html.EscapeString(comment.Content)
	_, execErr := prp.Exec(
		comment.ID,
		comment.UserID,
		comment.PostID,
		comment.Content,
		comment.CreatedAt,
	)
	if execErr != nil {
		return execErr
	}
	return nil
}

func (c *CommentRepositorie) GetCommentByPost(postID string, pagination int) ([]models.CommentDetails, error) {
	querySelect := `
	SELECT
	 comments.id AS comment_id,
	 comments.content,
	 comments.created_at,
	 users.id AS user_id,
	 users.username,
	 (SELECT COUNT(*) FROM likes WHERE likes.comment_id = comments.id AND likes.react_type = 'like') AS LikeCount,
	 (SELECT COUNT(*) FROM likes WHERE likes.comment_id = comments.id AND likes.react_type = 'dislike') AS DisLikeCount
	FROM 
	 comments
	JOIN
	 users ON comments.user_id = users.id
	WHERE
	 comments.post_id = ?
	ORDER BY
	 comments.created_at DESC
	 LIMIT 5 OFFSET ?;`

	rows, queryErr := c.DB.Query(querySelect, postID, pagination)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()
	var comments []models.CommentDetails
	for rows.Next() {
		var currentComment models.CommentDetails
		scanErr := rows.Scan(
			&currentComment.CommentID,
			&currentComment.Content,
			&currentComment.CreatedAt,
			&currentComment.UserID,
			&currentComment.Username,
			&currentComment.LikeCountComment,
			&currentComment.DisLikeCountComment,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		currentComment.FormattedDate = currentComment.CreatedAt.Format("01/02/2006, 3:04:05 PM")
		comments = append(comments, currentComment)
	}
	return comments, nil
}
