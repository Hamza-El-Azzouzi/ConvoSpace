package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Comment struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	PostID    string `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// type CommentWithUser struct {
// 	CommentID   uuid.UUID `json:"comment_id"`
// 	Content     string    `json:"content"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UserID      uuid.UUID `json:"user_id"`
// 	Username    string    `json:"username"`
// 	Email       string    `json:"email"`
// }