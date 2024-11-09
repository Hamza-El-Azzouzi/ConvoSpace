package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/internal/database"
	"forum/internal/handlers"
	"forum/internal/middleware"
	"forum/internal/repositories"
	"forum/internal/services"
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
	cleaner := &utils.Cleaner{Db:db}

	go cleaner.CleanupExpiredSessions()

	userRepo := &repositories.UserRepository{DB: db}
	categorieRepo := &repositories.CategoryRepository{DB: db}
	postRepo := &repositories.PostRepository{DB: db}
	commentRepo := &repositories.CommentRepositorie{DB: db}
	likeRepo := &repositories.LikeReposetorie{DB: db}

	postServices := &services.PostService{PostRepo: postRepo}
	categorieServices := &services.CategoryService{CategorieRepo: categorieRepo}
	commentService := &services.CommentService{CommentRepo: commentRepo}
	likeService := &services.LikeService{LikeRepo: likeRepo}
	authService := &services.AuthService{UserRepo: userRepo}

	authMidlware := &middleware.AuthMidlaware{AuthService: authService}

	authHandler := &handlers.AuthHandler{AuthService: authService, AuthMidlaware: authMidlware}
	postHandler := &handlers.PostHandler{
		AuthService:     authService,
		AuthMidlaware: authMidlware,
		CategoryService: categorieServices,
		PostService:     postServices,
		CommentService:  commentService,
		AuthHandler:     authHandler,
	}
	likeHandler := &handlers.LikeHandler{
		LikeService: likeService,
		AuthService: authService,
		AuthMidlaware: authMidlware,
	}

	mux := http.NewServeMux()

	fmt.Println("Starting the forum server...")
	mux.HandleFunc("/static/", utils.SetupStaticFilesHandlers)
	mux.HandleFunc("/", postHandler.HomeHandle)
	mux.HandleFunc("/create", postHandler.PostCreation)
	mux.HandleFunc("/createPost", postHandler.PostSaver)
	mux.HandleFunc("/sendcomment/", postHandler.CommentSaver)

	mux.HandleFunc("/logout", authHandler.LogoutHandle)
	mux.HandleFunc("/login", authHandler.LoginHandle)
	mux.HandleFunc("/register", authHandler.RegisterHandle)
	mux.HandleFunc("/detailsPost/", postHandler.DetailsPost)

	mux.HandleFunc("/like/", likeHandler.LikePost)
	mux.HandleFunc("/dislike/", likeHandler.DisLikePost)

	mux.HandleFunc("/likeComment/", likeHandler.LikeComment)
	mux.HandleFunc("/dislikeComment/", likeHandler.DisLikeComment)

	mux.HandleFunc("/filters", postHandler.PostFilter)

	mux.HandleFunc("/checker", authHandler.CheckDoubleLogging)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler, pattern := mux.Handler(r)
		if pattern == "" || pattern == "/" && r.URL.Path != "/" {
			utils.Error(w, 404)
			return
		}
		handler.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
