package services

import (
	"errors"
	"time"

	"forum/internal/models"
	"forum/internal/repositories"

	"github.com/gofrs/uuid/v5"
)

type CommentService struct {
	CommentRepo *repositories.CommentRepositorie
}

var (
	NotFoundPostId = errors.New("not found post id")
	EmptyComment   = errors.New("empty comment")
)

// get the data from the handler and create a ID with uuid then
// call the method from the repository
func (c *CommentService) SaveComment(userID uuid.UUID, postID, content string) error {
	comment := &models.Comment{
		ID:        uuid.Must(uuid.NewV4()),
		UserID:    userID,
		PostID:    postID,
		Content:   content,
		CreatedAt: time.Now().UTC(),
	}
	return c.CommentRepo.Create(comment)
}

// get all comments about a post that just got a new comment by calling the methode in repository
func (c *CommentService) GetCommentByPost(postID string) ([]models.CommentDetails, error) {
	if postID == "" {
		return nil, NotFoundPostId
	}

	comments, err := c.CommentRepo.GetCommentByPost(postID)
	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, EmptyComment
	}

	return comments, nil
}
