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
	var id *string
	if liked == "post" {
		id = like.PostID
	} else {
		id = like.CommentID
	}
	var reaction string
	row := l.DB.QueryRow("SELECT react_type FROM likes WHERE "+liked+"_id  AND like.UserID = ?",
		id, like.UserID)
	switch err := row.Scan(&reaction); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return nil
	case nil:
		if reaction == like.ReactType {
			_, err := l.DB.Exec("DELETE FROM likes WHERE id = ? AND user_id = ? ", id, like.UserID)
			if err != nil {
				return err
			}
			return nil
		} else {
			_, err := l.DB.Exec("UPDATE FROM likes WHERE id = ? AND user_id = ? ", id, like.UserID)
			if err != nil {
				return err
			}
			return nil
		}
	}
	_, err := l.DB.Exec("INSERT INTO likes (id, user_id, post_id, comment_id, react_type)VALUES (?, ?, ?, ?, ?)",
		like.ID, like.UserID, like.PostID, like.CommentID, like.ReactType,
	)
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
		return errcountlike, nil
	}
	errcountdislike := l.DB.QueryRow("SELECT COUNT(*) FROM likes WHERE "+liked+"_id = ? AND react_type = 'dislike'", Id).Scan(&dislike)
	if errcountdislike != nil {
		return errcountdislike, nil
	}
	data := map[string]any{
		"id":      Id,
		"like":    like,
		"dislike": dislike,
	}
	fmt.Println("dd", data)
	return data, nil
}
