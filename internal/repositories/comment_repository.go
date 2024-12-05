package repositories

import (
	"database/sql"

	"forum/internal/models"
)

type CommentRepositorie struct {
	DB *sql.DB
}
// save the comment in the data base and return error 
func (c *CommentRepositorie) Create(comment *models.Comment) error {
	
}


// Execute the SQL query to retrieve comments and their associated data
// The query fetches comment details, user details, and counts of likes and dislikes
// Iterate through the rows returned by the query
// Map the row data to the CommentDetails struct
// Format the comment creation date into a user-friendly string "01/02/2006, 3:04:05 PM"
func (c *CommentRepositorie) GetCommentByPost(postID string) ([]models.CommentDetails, error) {
	
}

