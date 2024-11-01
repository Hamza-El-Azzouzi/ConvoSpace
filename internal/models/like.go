package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Like struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	PostID    *string   `json:"post_id,omitempty"`
	CommentID *string   `json:"comment_id,omitempty"`
	ReactType string    `json:"react_type"`
	CreatedAt time.Time `json:"created_at"`
}
