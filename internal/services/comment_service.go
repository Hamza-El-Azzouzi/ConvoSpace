package services

import (
	"forum/internal/repositories"
)

type CommentService struct {
	CommentRepo *repositories.CommentRepositorie
}

// get the data from the handler and create a ID with uuid then
// call the method from the repository
// func (c *CommentService) SaveComment(userID uuid.UUID, postID, content string) error {

// }

// get all comments about a post that just got a new comment by calling the methode in repository
// func (c *CommentService) GetCommentByPost(postID string) ([]models.CommentDetails,error) {

// }
