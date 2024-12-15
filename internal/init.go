package internal

import (
	"database/sql"

	"forum/internal/handlers"
	"forum/internal/middleware"
	"forum/internal/repositories"
	"forum/internal/services"
)

func InitRepositories(db *sql.DB) (*repositories.UserRepository,
	*repositories.CategoryRepository,
	*repositories.PostRepository,
	*repositories.CommentRepositorie,
	*repositories.LikeReposetorie,
	*repositories.SessionsRepositorie,
) {
	return &repositories.UserRepository{DB: db},
		&repositories.CategoryRepository{DB: db},
		&repositories.PostRepository{DB: db},
		&repositories.CommentRepositorie{DB: db},
		&repositories.LikeReposetorie{DB: db},
		&repositories.SessionsRepositorie{DB: db}
}

func InitServices(userRepo *repositories.UserRepository,
	postRepo *repositories.PostRepository,
	categoryRepo *repositories.CategoryRepository,
	commentRepo *repositories.CommentRepositorie,
	likeRepo *repositories.LikeReposetorie,
	sessionRepo *repositories.SessionsRepositorie) (*services.AuthService,
	*services.PostService,
	*services.CategoryService,
	*services.CommentService,
	*services.LikeService,
	*services.SessionService,
) {
	return &services.AuthService{UserRepo: userRepo},
		&services.PostService{PostRepo: postRepo,CategoryRepo:categoryRepo},
		&services.CategoryService{CategorieRepo: categoryRepo},
		&services.CommentService{CommentRepo: commentRepo},
		&services.LikeService{LikeRepo: likeRepo},
		&services.SessionService{SessionRepo: sessionRepo}
}
// seesion repo
func InitHandlers(authService *services.AuthService,
	postService *services.PostService,
	categoryService *services.CategoryService,
	commentService *services.CommentService,
	likeService *services.LikeService,
	sessionRepo *services.SessionService,
	authMiddleware *middleware.AuthMiddleware) (*handlers.AuthHandler,
	*handlers.PostHandler,
	*handlers.LikeHandler,
) {
	authHandler := &handlers.AuthHandler{
		AuthService:    authService,
		AuthMidlaware:  authMiddleware,
		SessionService: sessionRepo,
	}
	postHandler := &handlers.PostHandler{
		AuthService:     authService,
		AuthMidlaware:   authMiddleware,
		CategoryService: categoryService,
		PostService:     postService,
		CommentService:  commentService,
		AuthHandler:     authHandler,
	}
	likeHandler := &handlers.LikeHandler{
		LikeService:   likeService,
		AuthMidlaware: authMiddleware,
	}

	return authHandler, postHandler, likeHandler
}
