package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Like struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	PostID    uuid.UUID `json:"post_id,omitempty"`
	CommentID uuid.UUID `json:"comment_id,omitempty"`
	IsLike    bool      `json:"is_like"` // true for like, false for dislike
	CreatedAt time.Time `json:"created_at"`
}
