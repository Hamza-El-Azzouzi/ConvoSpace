package repositories

import (
	"database/sql"

	"forum/internal/models"

	"github.com/gofrs/uuid/v5"
)

type LikeReposetorie struct {
	DB *sql.DB
}

func (l *LikeReposetorie) CreateLike(like *models.Like, liked string) error {
	var id *string
	var existingreactionID uuid.UUID
	var reaction string
	if liked == "post" {
		id = like.PostID
	} else {
		id = like.CommentID
	}
	row := l.DB.QueryRow("SELECT id, react_type FROM likes WHERE "+liked+"_id = ?  AND user_id = ?", id, like.UserID)
	switch err := row.Scan(&existingreactionID, &reaction); err {
	case sql.ErrNoRows:
	case nil:
		if reaction == like.ReactType {
			preparedQuery, err := l.DB.Prepare("DELETE FROM likes WHERE id = ? AND user_id = ? ")
			if err != nil {
				return err
			}
			_, err = preparedQuery.Exec(existingreactionID, like.UserID)
			if err != nil {
				return err
			}
			return nil
		} else {
			preparedQuery, err := l.DB.Prepare("UPDATE likes SET react_type = ? WHERE id = ?")
			if err != nil {
				return err
			}
			_, err = preparedQuery.Exec(like.ReactType, existingreactionID)
			if err != nil {
				return err
			}
			return nil
		}
	default:
		return err
	}
	preparedQuery, err := l.DB.Prepare("INSERT INTO likes (id, user_id, post_id, comment_id, react_type) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = preparedQuery.Exec(like.ID, like.UserID, like.PostID, like.CommentID, like.ReactType)
	if err != nil {
		return err
	}

	return nil
}

func (l *LikeReposetorie) GetLikes(Id, liked string) (any, error) {
	var like int
	var dislike int
	errcountlike := l.DB.QueryRow("SELECT COUNT(*) FROM likes WHERE "+liked+"_id = ? AND react_type = 'like'", Id).Scan(&like)
	if errcountlike != nil {
		return nil, errcountlike
	}
	errcountdislike := l.DB.QueryRow("SELECT COUNT(*) FROM likes WHERE "+liked+"_id = ? AND react_type = 'dislike'", Id).Scan(&dislike)
	if errcountdislike != nil {
		return nil, errcountdislike
	}

	data := map[string]any{
		"id":      Id,
		"like":    like,
		"dislike": dislike,
	}
	return data, nil
}
