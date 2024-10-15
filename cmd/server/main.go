package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/internal/handlers"
	"forum/internal/utils"
	"forum/internal/database"
    "forum/internal/repositories"
    "forum/internal/services"
)

func main() {
	db, err := database.InitDB("forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = database.RunMigrations(db)
	if err != nil {
		fmt.Printf("Error running migrations: %v", err)
	}

	userRepo := &repositories.UserRepository{DB: db}
	authService := &services.AuthService{UserRepo: userRepo}
	authHandler := &handlers.AuthHandler{AuthService: authService}

	fmt.Println("Starting the forum server...")
	http.HandleFunc("/static/", utils.SetupStaticFilesHandlers)
	http.HandleFunc("/", handlers.HomeHandle)
	http.HandleFunc("/logout", authHandler.LogoutHandle)
	http.HandleFunc("/login", authHandler.LoginHandle)
	http.HandleFunc("/register", authHandler.RegisterHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}