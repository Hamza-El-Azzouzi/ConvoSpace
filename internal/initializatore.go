package internal

import (
	"database/sql"

	"forum/internal/handlers"
	"forum/internal/middleware"
	"forum/internal/repositories"
	"forum/internal/services"
)

func InitRepositories(db *sql.DB) (*repositories.UserRepository, *repositories.CategoryRepository, *repositories.PostRepository, *repositories.CommentRepositorie, *repositories.LikeReposetorie) {
	return &repositories.UserRepository{DB: db},
		&repositories.CategoryRepository{DB: db},
		&repositories.PostRepository{DB: db},
		&repositories.CommentRepositorie{DB: db},
		&repositories.LikeReposetorie{DB: db}
}

func InitServices(userRepo *repositories.UserRepository, postRepo *repositories.PostRepository, categoryRepo *repositories.CategoryRepository, commentRepo *repositories.CommentRepositorie, likeRepo *repositories.LikeReposetorie) (*services.AuthService, *services.PostService, *services.CategoryService, *services.CommentService, *services.LikeService) {
	return &services.AuthService{UserRepo: userRepo},
		&services.PostService{PostRepo: postRepo},
		&services.CategoryService{CategorieRepo: categoryRepo},
		&services.CommentService{CommentRepo: commentRepo},
		&services.LikeService{LikeRepo: likeRepo}
}

func InitHandlers(authService *services.AuthService, postService *services.PostService, categoryService *services.CategoryService, commentService *services.CommentService, likeService *services.LikeService) (*handlers.AuthHandler, *handlers.PostHandler, *handlers.LikeHandler) {
	authMiddleware := &middleware.AuthMidlaware{AuthService: authService}

	authHandler := &handlers.AuthHandler{AuthService: authService, AuthMidlaware: authMiddleware}
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
		AuthService:   authService,
		AuthMidlaware: authMiddleware,
	}

	return authHandler, postHandler, likeHandler
}
