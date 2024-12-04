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
		log.Fatalf("error in DB : %v", err)
		return
	}
	err = database.RunMigrations(db)
	if err != nil {
		fmt.Printf("Error running migrations: %v", err)
		return
	}

	err = database.InsertDefaultCategories(db)
	if err != nil {
		fmt.Printf("error inserting default categories: %v", err)
		return
	}

	defer db.Close()

	userRepo, categoryRepo, postRepo, commentRepo, likeRepo, sessionRepo := internal.InitRepositories(db)

	authService, postService, categoryService, commentService, likeService, sessionService := internal.InitServices(userRepo,
		postRepo,
		categoryRepo,
		commentRepo,
		likeRepo,
		sessionRepo)

	authMiddleware := &middleware.AuthMiddleware{AuthService: authService,SessionService: sessionService}

	authHandler, postHandler, likeHandler := internal.InitHandlers(authService,
		postService,
		categoryService,
		commentService,
		likeService,
		sessionService,
		authMiddleware)

	cleaner := &utils.Cleaner{SessionService: sessionService}

	go cleaner.CleanupExpiredSessions()

	mux := http.NewServeMux()

	routes.SetupRoutes(mux, authHandler, postHandler, likeHandler, authMiddleware)

	fmt.Println("Starting the forum server...\nWelcome http://localhost:8082/")

	log.Fatal(http.ListenAndServe("0.0.0.0:8082", mux))
}
