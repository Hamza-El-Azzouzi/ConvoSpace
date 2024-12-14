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

func (c *CommentService) GetCommentByPost(postID string, pagination int) ([]models.CommentDetails, error) {
	comment, err := c.CommentRepo.GetCommentByPost(postID, pagination)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
