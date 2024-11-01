package services

import (
	"forum/internal/models"
	"forum/internal/repositories"

	"github.com/gofrs/uuid/v5"
)

type LikeService struct {
	LikeRepo *repositories.LikeReposetorie
}

func (l *LikeService) Create(userID uuid.UUID, postID, commentID, reactType string , isComment bool) error {
	likeID := uuid.Must(uuid.NewV4())
	var postIDPtr, commentIDPtr *string
	if postID != "" {
		postIDPtr = &postID
	}
	if commentID != "" {
		commentIDPtr = &commentID
	}
	like := &models.Like{
		ID:        likeID,
		UserID:    userID,
		PostID:    postIDPtr,
		CommentID: commentIDPtr,
		ReactType: reactType,
	}
	if isComment{
		return l.LikeRepo.CreateForComment(like)
	}
	return l.LikeRepo.CreateForPost(like)
}