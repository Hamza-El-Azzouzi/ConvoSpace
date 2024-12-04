package services

import (
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

func (l *LikeService) Create(userID uuid.UUID, postID, commentID, reactType string, isComment bool) error {

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

	if isComment {
		return l.LikeRepo.CreateLike(like, "comment")
	}

	return l.LikeRepo.CreateLike(like, "post")
}

// var reaction int
// pre, err := db.Prepare("SELECT is_like FROM likes WHERE post_id = ? AND user_id = ? ")
// if err != nil {
// 	fmt.Println(err.Error())
// 	return err
// }
// defer pre.Close()
// err = pre.QueryRow(id, id_user).Scan(&reaction)
// if err != nil {
// 	fmt.Println(err.Error())
// 	return err
// }
// if _, err := db.Exec(`DELETE FROM Likes WHERE post_id = ? AND user_id = ? `, id, id_user); err != nil {
// 	fmt.Println(err.Error())
// 	return err
// }
// if reaction == 2 {
// 	if _, err = db.Exec(`INSERT INTO likes (post_id,user_id,is_like) VALUES (?,?,?)`, id, id_user, 1); err != nil {
// 		fmt.Println(err.Error())
// 		return err
// 	}
// }
