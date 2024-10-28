package repositories

import (
	"database/sql"
	"fmt"

	"forum/internal/models"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}

// GetAllCategories gets raw data from database
func (r *PostRepository) Create(post *models.Post) error {
	_, err := r.DB.Exec(
		"INSERT INTO posts (ID, user_id, Title, Content) VALUES (?, ?, ?, ?)",
		post.ID, post.UserID, post.Title, post.Content,
	)
	return err
}

func (r *PostRepository) PostCatgorie(postCategorie *models.PostCategory) error {
	_, err := r.DB.Exec(
		"INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)",
		postCategorie.PostID, postCategorie.CategoryID,
	)
	return err
}

func (r *PostRepository) AllPosts() ([]models.PostWithUser, error) {
	query := `SELECT 
		posts.id AS post_id,
		posts.title,
		posts.content,
		posts.created_at,
		users.id AS user_id,
		users.username,
		users.email
	FROM 
		posts
	JOIN 
		users ON posts.user_id = users.id
	ORDER BY 
		posts.created_at DESC;`
	rows, err := r.DB.Query(query)
	if err != nil {
        return nil, fmt.Errorf("error querying posts with user info: %v", err)
    }
    defer rows.Close()

    var posts []models.PostWithUser
    for rows.Next() {
        var post models.PostWithUser
        if err := rows.Scan(
            &post.PostID,
            &post.Title,
            &post.Content,
            &post.CreatedAt,
            &post.UserID,
            &post.Username,
            &post.Email,
        ); err != nil {
            return nil, fmt.Errorf("error scanning post with user info: %v", err)
        }
		post.FormattedDate = post.CreatedAt.Format("January 2, 2006")
        posts = append(posts, post)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating posts with user info: %v", err)
    }

    return posts, nil
}
