package repositories

import (
	"database/sql"
	"fmt"
	"html"
	"strings"
	"time"

	"forum/internal/models"

	"github.com/gofrs/uuid/v5"
)

type PostRepository struct {
	DB *sql.DB
}

func (r *PostRepository) Create(post *models.Post) error {
	post.Content = html.EscapeString(post.Content)
	post.Title = html.EscapeString(post.Title)
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

func (r *PostRepository) AllPosts(pagination int) ([]models.PostWithUser, error) {
	query := `SELECT 
		posts.id AS post_id,
		posts.title,
		posts.content,
		posts.created_at,
		users.id AS user_id,
		users.username,
		REPLACE(IFNULL(GROUP_CONCAT(DISTINCT categories.name), ''), ',', ' | ') AS category_names,
		(SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) AS comment_count,
		(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id AND likes.react_type = "like") AS likes_count,
		(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id AND likes.react_type = "dislike") AS dislike_count,
		total_posts.total_count
		FROM 
			posts
		JOIN 
			users ON posts.user_id = users.id
		LEFT JOIN 
			post_categories ON posts.id = post_categories.post_id
		LEFT JOIN 
			categories ON post_categories.category_id = categories.id
		CROSS JOIN 
    		(SELECT COUNT(*) AS total_count FROM posts) AS total_posts
		GROUP BY 
			posts.id
		ORDER BY posts.created_at DESC LIMIT 5 OFFSET ?;`
	rows, err := r.DB.Query(query, pagination)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error querying posts with user info: %v", err)
	}
	defer rows.Close()

	var posts []models.PostWithUser
	for rows.Next() {
		var post models.PostWithUser
		err = rows.Scan(
			&post.PostID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UserID,
			&post.Username,
			&post.CategoryName,
			&post.CommentCount,
			&post.LikeCount,
			&post.DisLikeCount,
			&post.TotalCount,
		)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("error scanning post with user info: %v", err)
		}
		post.FormattedDate = post.CreatedAt.Format("01/02/2006, 3:04:05 PM")
		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
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
	    REPLACE(IFNULL(GROUP_CONCAT(DISTINCT categories.name), ''), ',', ' | ') AS category_names,

		(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id AND likes.react_type = "like") AS likes_count,
		(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id AND likes.react_type = "dislike") AS dislike_count,
	    comments.id AS comment_id,
	    comments.content AS comment_content,
	    comments.created_at AS comment_created_at,
	    comment_user.id AS comment_user_id,
	    comment_user.username AS comment_username,
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
			comments.id
		ORDER BY 
			comments.created_at DESC;`

	rows, err := r.DB.Query(query, PostId)
	if err != nil {
		return models.PostDetails{}, fmt.Errorf("error f query : %v", err)
	}
	defer rows.Close()

	var postDetails models.PostDetails
	postDetails.Comments = []models.CommentDetails{}

	for rows.Next() {
		var (
			postID              uuid.UUID
			title               string
			content             string
			createdAt           time.Time
			userID              uuid.UUID
			username            string
			categoryNames       string
			likeCount           int
			disLikeCount        int
			commentID           sql.NullString
			commentContent      sql.NullString
			commentCreated      sql.NullTime
			commentUserID       sql.NullString
			commentUsername     sql.NullString
			commentLikesCount   sql.NullInt64
			commentDislikeCount sql.NullInt64
		)

		err = rows.Scan(
			&postID,
			&title,
			&content,
			&createdAt,
			&userID,
			&username,
			&categoryNames,
			&likeCount,
			&disLikeCount,
			&commentID,
			&commentContent,
			&commentCreated,
			&commentUserID,
			&commentUsername,
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
				FormattedDate: createdAt.Format("01/02/2006, 3:04:05 PM"),
				CategoryNames: categoryNames,
				LikeCount:     likeCount,
				DisLikeCount:  disLikeCount,
			}
		}

		if commentID.Valid {
			parsedCommentID, err := uuid.FromString(commentID.String)
			if err != nil {
				return models.PostDetails{}, fmt.Errorf("error f parse com id : %v", err)
			}

			parsedUserIDComment, err := uuid.FromString(commentUserID.String)
			if err != nil {
				return models.PostDetails{}, fmt.Errorf("error f parse : %v", err)
			}
			comment := models.CommentDetails{
				CommentID:           parsedCommentID,
				Content:             commentContent.String,
				CreatedAt:           commentCreated.Time,
				UserID:              parsedUserIDComment,
				Username:            commentUsername.String,
				FormattedDate:       commentCreated.Time.Format("01/02/2006, 3:04:05 PM"),
				LikeCountComment:    commentLikesCount.Int64,
				DisLikeCountComment: commentDislikeCount.Int64,
			}
			postDetails.Comments = append(postDetails.Comments, comment)
		}
	}
	err = rows.Err()
	if err != nil {
		return models.PostDetails{}, fmt.Errorf("error f row.errr : %v", err)
	}
	return postDetails, nil
}

func (r *PostRepository) FilterPost(filterby, category string, userID uuid.UUID, pagination int) ([]models.PostWithUser, error) {
	// Base Query
	baseQuery := `
		SELECT 
			posts.id AS post_id,
			posts.title,
			posts.content,
			posts.created_at,
			users.id AS user_id,
			users.username,
			REPLACE(IFNULL(GROUP_CONCAT(DISTINCT categories.name), ''), ',', ' | ') AS category_names,
			(SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) AS comment_count,
			(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id AND likes.react_type = "like") AS likes_count,
			(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id AND likes.react_type = "dislike") AS dislike_count,
			COUNT(*) OVER() AS total_count
		FROM 
			posts
		JOIN 
			users ON posts.user_id = users.id
		LEFT JOIN 
			post_categories ON posts.id = post_categories.post_id
		LEFT JOIN 
			categories ON post_categories.category_id = categories.id
		LEFT JOIN 
			likes ON posts.id = likes.post_id`

	// Query Components
	var whereClauses []string
	args := []any{}
	if filterby == "created" {
		whereClauses = append(whereClauses, "posts.user_id = ?")
		args = append(args, userID)
	} else if filterby == "liked" {
		whereClauses = append(whereClauses, "likes.user_id = ? AND likes.comment_id IS NULL")
		args = append(args, userID)
	}

	if category != "" {
		whereClauses = append(whereClauses, "post_categories.category_id = ?")
		args = append(args, category)
	}

	whereQuery := ""
	if len(whereClauses) > 0 {
		whereQuery = " WHERE " + strings.Join(whereClauses, " AND ")
	}

	groupQuery := " GROUP BY posts.id"
	orderQuery := " ORDER BY posts.created_at DESC"
	limitQuery := " LIMIT 5 OFFSET ?"
	args = append(args, pagination)

	finalQuery := baseQuery + whereQuery + groupQuery + orderQuery + limitQuery

	rows, err := r.DB.Query(finalQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var posts []models.PostWithUser
	for rows.Next() {
		var post models.PostWithUser
		err = rows.Scan(
			&post.PostID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UserID,
			&post.Username,
			&post.CategoryName,
			&post.CommentCount,
			&post.LikeCount,
			&post.DisLikeCount,
			&post.TotalCount,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning post data: %w", err)
		}
		post.FormattedDate = post.CreatedAt.Format("01/02/2006, 3:04:05 PM")
		posts = append(posts, post)
	}

	return posts, nil
}
