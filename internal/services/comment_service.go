package services

import (
	"forum/internal/models"
	"forum/internal/repositories"

	"github.com/gofrs/uuid/v5"
)

type CommentService struct {
	CommentRepo *repositories.CommentRepositorie
}

func (c *CommentService) SaveComment(userID uuid.UUID, postID, content string) error {
	commentID := uuid.Must(uuid.NewV4())
	comment := &models.Comment{
		ID: commentID,
		UserID: userID,
		PostID: postID,
		Content: content,
	}
	return c.CommentRepo.Create(comment)
}
