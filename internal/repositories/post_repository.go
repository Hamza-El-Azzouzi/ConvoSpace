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
    posts.id;`
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
		(SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) AS comment_count,
		(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id AND likes.react_type = "like") AS likes_count,
		(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id AND likes.react_type = "dislike") AS dislike_count,
	    comments.id AS comment_id,
		comments.post_id as post_id,
	    comments.content AS comment_content,
	    comments.created_at AS comment_created_at,
	    comment_user.id AS comment_user_id,
	    comment_user.username AS comment_username,
	    comment_user.email AS comment_email,
		(SELECT COUNT(*) FROM likes WHERE likes.comment_id = comments.id AND likes.react_type = "like") AS comment_likes_count,
	    (SELECT COUNT(*) FROM likes WHERE likes.comment_id = comments.id AND likes.react_type = "dislike") AS comment_dislike_count
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
		return models.PostDetails{}, fmt.Errorf("error f query : %v", err)
	}
	defer rows.Close()

	var postDetails models.PostDetails
	postDetails.Comments = []models.CommentDetails{} // Initialize slice for comments

	for rows.Next() {
		var (
			postID              uuid.UUID
			title               string
			content             string
			createdAt           time.Time
			userID              uuid.UUID
			username            string
			email               string
			categoryNames       string
			commentCount        int
			likeCount           int
			disLikeCount        int
			commentID           sql.NullString
			postIDcomment       sql.NullString
			commentContent      sql.NullString
			commentCreated      sql.NullTime
			commentUserID       sql.NullString
			commentUsername     sql.NullString
			commentEmail        sql.NullString
			commentLikesCount   sql.NullInt64
			commentDislikeCount sql.NullInt64
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
			&commentCount,
			&likeCount,
			&disLikeCount,
			&commentID,
			&postIDcomment,
			&commentContent,
			&commentCreated,
			&commentUserID,
			&commentUsername,
			&commentEmail,
			&commentLikesCount,
			&commentDislikeCount,
		)
		if err != nil {
			return models.PostDetails{}, fmt.Errorf("error f row scan : %v", err)
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
				CommentCount:  commentCount,
				LikeCount:     likeCount,
				DisLikeCount:  disLikeCount,
			}
		}

		if commentID.Valid {
			parsedCommentID, err := uuid.FromString(commentID.String)
			if err != nil {
				return models.PostDetails{}, fmt.Errorf("error f parse com id : %v", err)
			}
			parsedposttID, err := uuid.FromString(postIDcomment.String)
			if err != nil {
				return models.PostDetails{}, fmt.Errorf("error f parse com id : %v", err)
			}
			parsedUserIDComment, err := uuid.FromString(commentUserID.String)
			if err != nil {
				return models.PostDetails{}, fmt.Errorf("error f parse : %v", err)
			}
			comment := models.CommentDetails{
				CommentID:           parsedCommentID,
				PostIDcomment:       parsedposttID,
				Content:             commentContent.String,
				CreatedAt:           commentCreated.Time,
				UserID:              parsedUserIDComment,
				Username:            commentUsername.String,
				Email:               commentEmail.String,
				FormattedDate:       commentCreated.Time.Format("January 2, 2006"),
				LikeCountComment:    commentLikesCount.Int64,
				DisLikeCountComment: commentDislikeCount.Int64,
			}
			postDetails.Comments = append(postDetails.Comments, comment)
		}
	}

	if err = rows.Err(); err != nil {
		return models.PostDetails{}, fmt.Errorf("error f row.errr : %v", err)
	}
	return postDetails, nil
}

func (r *PostRepository) FilterPost(post, like string, categorie []string) ([]models.PostWithUser, error) {
	baseQuery := `SELECT 
    posts.id AS post_id,
    posts.title,
    posts.content,
    posts.created_at,
    users.id AS user_id,
    users.username,
    users.email,
	(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id AND react_type = "like") AS like_count,
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
    	comments ON posts.id = comments.post_id`
	groupQuery := " GROUP BY posts.id"
	var categoryFilter string
	args := []interface{}{}
	fmt.Printf("slice f repo : %v  hada len : %v \n" ,categorie,len(categorie))
	if len(categorie) > 0 {
		categoryFilter = " WHERE post_categories.category_id IN ("
		for i, id := range categorie {
			if i > 0 {
				categoryFilter += ", "
			}
			categoryFilter += "?"
			args = append(args, id)
		}
		categoryFilter += ")"
	}
	orderClause := ""
	if post != "" || like != "" {
		orderClause = " ORDER BY"
	}
	if post == "asc" {
		orderClause += " posts.created_at ASC" 
	} else if post == "desc" {
		orderClause += " posts.created_at DESC"
	}
	if post != "" && like != ""{
		orderClause += ","
	}
	if like == "asc" {
		orderClause += " like_count ASC" 
	} else if like == "desc" {
		orderClause += " like_count DESC"
	}

	finalQuery := baseQuery + categoryFilter + groupQuery+orderClause
	fmt.Println(finalQuery)
	fmt.Println(args...)
	rows, err := r.DB.Query(finalQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query in filter: %w", err)
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
			&post.LikeCount,
			&post.CategoryName,
			&post.CommentCount,
		); err != nil {
			return nil, fmt.Errorf("error scanning post with user info filter: %v", err)
		}
		post.FormattedDate = post.CreatedAt.Format("January 2, 2006")
		posts = append(posts, post)
	}

	return posts, nil
}
