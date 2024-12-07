package services

import (
	"time"

	"forum/internal/models"
	"forum/internal/repositories"

	"github.com/gofrs/uuid/v5"
)

type CommentService struct {
	CommentRepo *repositories.CommentRepositorie
}

// get the data from the handler and create a ID with uuid then
// call the method from the repository
func (c *CommentService) SaveComment(userID uuid.UUID, postID, content string) error {
	commentID := uuid.Must(uuid.NewV4())
	comment := &models.Comment{
		ID:        commentID,
		UserID:    userID,
		PostID:    postID,
		Content:   content,
		CreatedAt: time.Now(),
	}
	return c.CommentRepo.Create(comment)
}

// get all comments about a post that just got a new comment by calling the methode in repository
// func (c *CommentService) GetCommentByPost(postID string) ([]models.CommentDetails,error) {

// }
