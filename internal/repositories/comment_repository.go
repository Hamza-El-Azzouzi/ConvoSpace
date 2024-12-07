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
	data, _ := c.DB.Prepare("INSERT INTO comments (id, user_id, post_id, content, created_at) VALUES (?, ?, ?, ?, ?)")
	defer data.Close()
	_, err := data.Exec(
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
// func (c *CommentRepositorie) GetCommentByPost(postID string) ([]models.CommentDetails, error) {

// }
