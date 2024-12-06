package repositories

import (
	"database/sql"
	"fmt"

	"forum/internal/models"
)

type LikeReposetorie struct {
	DB *sql.DB
}

func (l *LikeReposetorie) CreateLike(like *models.Like, liked string) error {

}






































func (l *LikeReposetorie) GetLikes(Id, liked string) (any, error) {
	var like int
	var dislike int
	errLike := l.DB.QueryRow(
		"SELECT COUNT(*) FROM likes WHERE "+liked+"_id = ? AND react_type = 'like'",
		Id,
	).Scan(&like)
	errDislike := l.DB.QueryRow(
		"SELECT COUNT(*) FROM likes WHERE "+liked+"_id = ? AND react_type = 'dislike'",
		Id,
	).Scan(&dislike)
	if errDislike != nil || errLike != nil {
		return nil, fmt.Errorf("error : %v , %v ", errDislike, errLike)
	}
	data := map[string]any{
		"id":           Id,
		"likeCount":    like,
		"dislikeCount": dislike,
	}
	return data, nil
}
