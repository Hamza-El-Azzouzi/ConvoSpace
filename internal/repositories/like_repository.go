package repositories

import (
	"database/sql"
	"fmt"

	"forum/internal/models"

	"github.com/gofrs/uuid/v5"
)

type LikeReposetorie struct {
	DB *sql.DB
}

func (l *LikeReposetorie) CreateLike(like *models.Like, liked string) error {
	var existingID uuid.UUID
	var existingReactType string
	var reactOn *string
	if liked == "post" {
		reactOn = like.PostID
	} else {
		reactOn = like.CommentID
	}
	err := l.DB.QueryRow(
		"SELECT id, react_type FROM likes WHERE user_id = ? AND "+liked+"_id = ?",
		like.UserID, reactOn,
	).Scan(&existingID, &existingReactType)

	if err == nil {
		if existingReactType == like.ReactType {

			_, err = l.DB.Exec("DELETE FROM likes WHERE id = ?", existingID)
			if err != nil {
				return fmt.Errorf("failed to remove existing reaction: %w", err)
			}
			return nil
		} else {
			_, err = l.DB.Exec(
				"UPDATE likes SET react_type = ? WHERE id = ?",
				like.ReactType, existingID,
			)
			if err != nil {
				return fmt.Errorf("failed to update reaction type: %w", err)
			}
			return nil
		}
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("failed to check existing reaction: %w", err)
	}
	_, err = l.DB.Exec(
		"INSERT INTO likes (id, user_id, post_id, comment_id, react_type) VALUES (?, ?, ?, ?, ?)",
		like.ID, like.UserID, like.PostID, like.CommentID, like.ReactType,
	)
	if err != nil {
		return fmt.Errorf("failed to create new reaction: %w", err)
	}
	return nil
}

func (l *LikeReposetorie) GetLikesPost(postID string) (any, error) {
	var like int
	var dislike int
	errLike := l.DB.QueryRow(
		"SELECT COUNT(*) FROM likes WHERE post_id = ? AND react_type = 'like'",
		postID,
	).Scan(&like)
	errDislike := l.DB.QueryRow(
		"SELECT COUNT(*) FROM likes WHERE post_id = ? AND react_type = 'dislike'",
		postID,
	).Scan(&dislike)
	if errDislike != nil || errLike != nil {
		return nil, fmt.Errorf("error : %v , %v ",errDislike,errLike)
	}
	data := map[string]any{
		"postID": postID,
		"likeCount":    like,
		"dislikeCount": dislike,
	}
	return data, nil
}
func (l *LikeReposetorie) GetLikesComment(commentID string) (any, error) {
	var like int
	var dislike int
	errLike := l.DB.QueryRow(
		"SELECT COUNT(*) FROM likes WHERE comment_id = ? AND react_type = 'like'",
		commentID,
	).Scan(&like)
	errDislike := l.DB.QueryRow(
		"SELECT COUNT(*) FROM likes WHERE comment_id = ? AND react_type = 'dislike'",
		commentID,
	).Scan(&dislike)
	if errDislike != nil || errLike != nil {
		return nil, fmt.Errorf("error : %v , %v ",errDislike,errLike)
	}
	data := map[string]any{
		"CommentID": commentID,
		"likeCount":    like,
		"dislikeCount": dislike,
	}
	return data, nil
}