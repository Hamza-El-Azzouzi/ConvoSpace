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
	comment.Content = html.EscapeString(comment.Content)
	_, err := c.DB.Exec("INSERT INTO comments (id, user_id, post_id, content, created_at) VALUES (?, ?, ?, ?, ?)",
		comment.ID,
		comment.UserID,
		comment.PostID,
		comment.Content,
		comment.CreatedAt,
	)
	return err
}

// Execute the SQL query to retrieve comments and their associated data
// The query fetches comment details, user details, and counts of likes and dislikes
// Iterate through the rows returned by the query
// Map the row data to the CommentDetails struct
// Format the comment creation date into a user-friendly string "01/02/2006, 3:04:05 PM"
func (c *CommentRepositorie) GetCommentByPost(postID string) ([]models.CommentDetails, error) {
	query := `
    SELECT 
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
		
		ORDER BY comments.created_at DESC;
		`
	var comments []models.CommentDetails
	rows, err := c.DB.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
