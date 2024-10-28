package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}


type PostCategory struct {
	PostID     uuid.UUID `json:"post_id"`
	CategoryID string `json:"category_id"`
}

type PostWithUser struct {
    PostID    uuid.UUID `json:"post_id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UserID    uuid.UUID `json:"user_id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
	FormattedDate string  `json:"formatted_date"` // Add this line
}