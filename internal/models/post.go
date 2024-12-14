package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Post struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Title     string
	Content   string
	CreatedAt time.Time
}

type PostCategory struct {
	PostID     uuid.UUID
	CategoryID string
}

type PostWithUser struct {
	PostID        uuid.UUID
	Title         string
	Content       string
	CreatedAt     time.Time
	UserID        uuid.UUID
	Username      string
	FormattedDate string
	CategoryName  string
	CommentCount  int
	LikeCount     int
	DisLikeCount  int
	TotalCount    string
}

type PostDetails struct {
	PostID        uuid.UUID
	Title         string
	Content       string
	CreatedAt     time.Time
	UserID        uuid.UUID
	Username      string
	FormattedDate string
	CategoryNames string
	CommentCount  int
	LikeCount     int
	DisLikeCount  int
	Comments      []CommentDetails
}

type CommentDetails struct {
	CommentID           uuid.UUID
	PostIDcomment       uuid.UUID
	Content             string
	CreatedAt           time.Time
	UserID              uuid.UUID
	Username            string
	FormattedDate       string
	LikeCountComment    int64
	DisLikeCountComment int64
}
