package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/internal/database"
	"forum/internal/handlers"
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
	userRepo := &repositories.UserRepository{DB: db}
	categorieRepo := &repositories.CategoryRepository{DB: db}
	postRepo := &repositories.PostRepository{DB: db}
	commentRepo := &repositories.CommentRepositorie{DB: db}

	postServices := &services.PostService{PostRepo: postRepo}
	categorieServices := &services.CategoryService{CategorieRepo: categorieRepo}
	commentService := &services.CommentService{CommentRepo: commentRepo}
	authService := &services.AuthService{UserRepo: userRepo}
	authHandler := &handlers.AuthHandler{AuthService: authService}
	postHandler := &handlers.PostHandler{
		AuthService:     authService,
		CategoryService: categorieServices,
		PostService:     postServices,
		CommentService:  commentService,
	}

	// postHandler := &handlers.PostHandler{AuthService: authService}

	fmt.Println("Starting the forum server...")
	http.HandleFunc("/static/", utils.SetupStaticFilesHandlers)
	http.HandleFunc("/", postHandler.HomeHandle)
	http.HandleFunc("/create", postHandler.PostCreation)
	http.HandleFunc("/createPost", postHandler.PostSaver)
	http.HandleFunc("/sendcomment/", postHandler.CommentSaver)

	http.HandleFunc("/logout", authHandler.LogoutHandle)
	http.HandleFunc("/login", authHandler.LoginHandle)
	http.HandleFunc("/register", authHandler.RegisterHandle)
	http.HandleFunc("/detailsPost/", postHandler.DetailsPost)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
