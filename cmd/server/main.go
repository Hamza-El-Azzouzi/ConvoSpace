package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/internal"
	"forum/internal/database"
	"forum/internal/middleware"
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

	if err := database.RunMigrations(db); err != nil {
		fmt.Printf("Error running migrations: %v", err)
	}

	if err := database.InsertDefaultCategories(db); err != nil {
		fmt.Printf("error inserting default categories: %v", err)
	}

	defer db.Close()

	userRepo, categoryRepo, postRepo, commentRepo, likeRepo, sessionRepo := internal.InitRepositories(db)

	authService, postService, categoryService, commentService, likeService, sessionService := internal.InitServices(userRepo, postRepo, categoryRepo, commentRepo, likeRepo, sessionRepo)
	authMiddleware := &middleware.AuthMiddleware{AuthService: authService}
	authHandler, postHandler, likeHandler := internal.InitHandlers(authService, postService, categoryService, commentService, likeService, sessionService, authMiddleware)
	mux := http.NewServeMux()
	cleaner := &utils.Cleaner{SessionService: sessionService}

	go cleaner.CleanupExpiredSessions()

	fmt.Println("Starting the forum server...\nWelcome http://localhost:8082/")

	routes.SetupRoutes(mux, authHandler, postHandler, likeHandler, authMiddleware)

	log.Fatal(http.ListenAndServe("0.0.0.0:8082", nil))
}
