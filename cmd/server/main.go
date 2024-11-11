package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/internal"
	"forum/internal/database"
	"forum/internal/routes"
	"forum/internal/utils"
)

func main() {
	db, err := database.InitDB("forum.db")
	if err != nil {
		log.Fatal(err)
		fmt.Printf("errr f DB : %v", err)
		return
	}
	err = database.RunMigrations(db)
	if err != nil {
		fmt.Printf("Error running migrations: %v", err)
	}
	if err := database.InsertDefaultCategories(db); err != nil {
		fmt.Printf("error inserting default categories: %v", err)
	}

	defer db.Close()
	cleaner := &utils.Cleaner{Db: db}

	go cleaner.CleanupExpiredSessions()

	// userRepo := &repositories.UserRepository{DB: db}
	// categorieRepo := &repositories.CategoryRepository{DB: db}
	// postRepo := &repositories.PostRepository{DB: db}
	// commentRepo := &repositories.CommentRepositorie{DB: db}
	// likeRepo := &repositories.LikeReposetorie{DB: db}

	// postServices := &services.PostService{PostRepo: postRepo}
	// categorieServices := &services.CategoryService{CategorieRepo: categorieRepo}
	// commentService := &services.CommentService{CommentRepo: commentRepo}
	// likeService := &services.LikeService{LikeRepo: likeRepo}
	// authService := &services.AuthService{UserRepo: userRepo}

	// authMidlware := &middleware.AuthMidlaware{AuthService: authService}

	// authHandler := &handlers.AuthHandler{AuthService: authService, AuthMidlaware: authMidlware}
	// postHandler := &handlers.PostHandler{
	// 	AuthService:     authService,
	// 	AuthMidlaware: authMidlware,
	// 	CategoryService: categorieServices,
	// 	PostService:     postServices,
	// 	CommentService:  commentService,
	// 	AuthHandler:     authHandler,
	// }
	// likeHandler := &handlers.LikeHandler{
	// 	LikeService: likeService,
	// 	AuthService: authService,
	// 	AuthMidlaware: authMidlware,
	// }

	userRepo, categoryRepo, postRepo, commentRepo, likeRepo := internal.InitRepositories(db)
	authService, postService, categoryService, commentService, likeService := internal.InitServices(userRepo, postRepo, categoryRepo, commentRepo, likeRepo)
	authHandler, postHandler, likeHandler := internal.InitHandlers(authService, postService, categoryService, commentService, likeService)
	mux := http.NewServeMux()

	fmt.Println("Starting the forum server...")

	routes.SetupRoutes(mux, authHandler, postHandler, likeHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
