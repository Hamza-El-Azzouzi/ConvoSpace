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
	CategoryID string    `json:"category_id"`
}

type PostWithUser struct {
	PostID        uuid.UUID `json:"post_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
	UserID        uuid.UUID `json:"user_id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	FormattedDate string    `json:"formatted_date"`
	CategoryName  string    // Add this line
	CommentCount int		`json:"comment_count"`
}

// PostDetails represents the details of a post including its categories and comments.

type PostDetails struct {
	PostID        uuid.UUID       `json:"post_id"`
	Title         string          `json:"title"`
	Content       string          `json:"content"`
	CreatedAt     time.Time       `json:"created_at"`
	UserID        uuid.UUID       `json:"user_id"`
	Username      string          `json:"username"`
	Email         string          `json:"email"`
	FormattedDate string          `json:"formatted_date"`
	CategoryNames string          `json:"category_names"`
	Comments      []CommentDetails `json:"comments"` // Slice of comments associated with the post
}

type CommentDetails struct {
	CommentID       uuid.UUID `json:"comment_id"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"created_at"`
	UserID          uuid.UUID `json:"user_id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	FormattedDate   string    `json:"formatted_date"`
}