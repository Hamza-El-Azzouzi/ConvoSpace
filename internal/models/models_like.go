package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)
type Like struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	PostID    *string
	CommentID *string
	ReactType string
	CreatedAt time.Time 
}
