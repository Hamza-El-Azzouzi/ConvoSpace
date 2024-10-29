package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"forum/internal/models"

	"github.com/gofrs/uuid/v5"
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
    users.email,
    IFNULL(GROUP_CONCAT(categories.name, ', '), '') AS category_names,
    COUNT(comments.id) AS comment_count
FROM 
    posts
JOIN 
    users ON posts.user_id = users.id
LEFT JOIN 
    post_categories ON posts.id = post_categories.post_id
LEFT JOIN 
    categories ON post_categories.category_id = categories.id
LEFT JOIN 
    comments ON posts.id = comments.post_id
GROUP BY 
    posts.id
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
			&post.CategoryName,
			&post.CommentCount,
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
func (r *PostRepository) GetPostById(PostId string) (models.PostDetails, error) {
	query := `SELECT 
	    posts.id AS post_id,
	    posts.title,
	    posts.content AS post_content,
	    posts.created_at AS post_created_at,
	    post_user.id AS post_user_id,
	    post_user.username AS post_username,
	    post_user.email AS post_email,
	    GROUP_CONCAT(DISTINCT categories.name) AS category_names,
	    comments.id AS comment_id,
	    comments.content AS comment_content,
	    comments.created_at AS comment_created_at,
	    comment_user.id AS comment_user_id,
	    comment_user.username AS comment_username,
	    comment_user.email AS comment_email
	FROM 
	    posts
	JOIN 
	    users AS post_user ON posts.user_id = post_user.id
	LEFT JOIN 
	    post_categories ON posts.id = post_categories.post_id
	LEFT JOIN 
	    categories ON post_categories.category_id = categories.id
	LEFT JOIN 
	    comments ON posts.id = comments.post_id
	LEFT JOIN 
	    users AS comment_user ON comments.user_id = comment_user.id
	WHERE 
	    posts.id = ?
	GROUP BY 
	    posts.id, comments.id
	ORDER BY 
	    posts.created_at DESC, comments.created_at ASC;`

	rows, err := r.DB.Query(query, PostId)
	if err != nil {
		return models.PostDetails{}, err
	}
	defer rows.Close()

	var postDetails models.PostDetails
	postDetails.Comments = []models.CommentDetails{} // Initialize slice for comments

	for rows.Next() {
		var (
			postID         uuid.UUID
			title          string
			content        string
			createdAt      time.Time
			userID         uuid.UUID
			username       string
			email          string
			categoryNames  string
			commentID      sql.NullString
			commentContent sql.NullString
			commentCreated sql.NullTime
			commentUserID  sql.NullString
			commentUsername sql.NullString
			commentEmail    sql.NullString
		)

		err := rows.Scan(
			&postID,
			&title,
			&content,
			&createdAt,
			&userID,
			&username,
			&email,
			&categoryNames,
			&commentID,
			&commentContent,
			&commentCreated,
			&commentUserID,
			&commentUsername,
			&commentEmail,
		)
		if err != nil {
			return models.PostDetails{}, err
		}

		if postDetails.PostID == uuid.Nil {
			postDetails = models.PostDetails{
				PostID:        postID,
				Title:         title,
				Content:       content,
				CreatedAt:     createdAt,
				UserID:        userID,
				Username:      username,
				Email:         email,
				FormattedDate: createdAt.Format("January 2, 2006"),
				CategoryNames: categoryNames,
			}
		}

		if commentID.Valid {
			parsedCommentID, err := uuid.FromString(commentID.String)
			if err != nil {
				return models.PostDetails{}, err
			}
			parsedUserIDComment, err := uuid.FromString(commentUserID.String)
			if err != nil {
				return models.PostDetails{}, err
			}
			comment := models.CommentDetails{
				CommentID:     parsedCommentID,
				Content:       commentContent.String,
				CreatedAt:     commentCreated.Time,
				UserID:        parsedUserIDComment,
				Username:      commentUsername.String,
				Email:         commentEmail.String,
				FormattedDate: commentCreated.Time.Format("January 2, 2006"),
			}
			postDetails.Comments = append(postDetails.Comments, comment)
		}
	}

	if err = rows.Err(); err != nil {
		return models.PostDetails{}, err
	}

	return postDetails, nil
}
