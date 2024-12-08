package services

import (
	"fmt"

	"forum/internal/models"
	"forum/internal/repositories"

	"github.com/gofrs/uuid/v5"
)

type LikeService struct {
	LikeRepo *repositories.LikeReposetorie
}

func (l *LikeService) GetLikes(ID, liked string) (any, error) {
	data, err := l.LikeRepo.GetLikes(ID, liked)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (l *LikeService) Create(userID uuid.UUID, postID, commentID string, typeOfReact, liked string) error {
	likeID, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("failed to generate UUID: %v", err)
	}

	var postIDPtr, commentIDPtr *string
	postIDPtr = &postID
	commentIDPtr = &commentID

	fmt.Println("l", postIDPtr)
	fmt.Println("f", commentID)
	fmt.Println("o", likeID)
	fmt.Println("m", typeOfReact)

	like := &models.Like{
		ID:        likeID,
		UserID:    userID,
		PostID:    postIDPtr,
		CommentID: commentIDPtr,
		ReactType: typeOfReact, // like or dislike
	}

	if liked == "post" {
		l.LikeRepo.CreateLike(like, "post")
	} else {
		l.LikeRepo.CreateLike(like, "comment")
	}
	return nil
}
