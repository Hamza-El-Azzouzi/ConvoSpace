package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"forum/internal/services"

	"github.com/gofrs/uuid/v5"
)

type LikeHandler struct {
	LikeService *services.LikeService
	AuthService *services.AuthService
}

func (l *LikeHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"LoggedIn": true,
	}
	var userID uuid.UUID
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
	postID := pathParts[2]
	cookie, err := r.Cookie("session_id")
	if err != nil {
		data["LoggedIn"] = false
	}

	if err == nil && cookie != nil {
		sessionId := cookie.Value
		user, err := l.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			userID = user.ID
			l.LikeService.Create(userID,postID,"","like")
			http.Redirect(w, r, fmt.Sprintf("http://localhost:8080/detailsPost/%v", postID), http.StatusSeeOther)
		} else {
			fmt.Printf("Error fetching user: %v", err)
		}
	}
}
func (l *LikeHandler) DisLikePost(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"LoggedIn": true,
	}
	var userID uuid.UUID
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
	postID := pathParts[2]
	cookie, err := r.Cookie("session_id")
	if err != nil {
		data["LoggedIn"] = false
	}

	if err == nil && cookie != nil {
		sessionId := cookie.Value
		user, err := l.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			userID = user.ID
			l.LikeService.Create(userID,postID,"","dislike")
			http.Redirect(w, r, fmt.Sprintf("http://localhost:8080/detailsPost/%v", postID), http.StatusSeeOther)
		} else {
			fmt.Printf("Error fetching user: %v", err)
		}
	}
}




func (l *LikeHandler) LikeComment(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"LoggedIn": true,
	}
	var userID uuid.UUID
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
	fmt.Println(pathParts)
	postID := pathParts[2]
	commentID := pathParts[3]
	cookie, err := r.Cookie("session_id")
	if err != nil {
		data["LoggedIn"] = false
	}

	if err == nil && cookie != nil {
		sessionId := cookie.Value
		user, err := l.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			userID = user.ID
			l.LikeService.Create(userID,postID,commentID,"like")
			http.Redirect(w, r, fmt.Sprintf("http://localhost:8080/detailsPost/%v", postID), http.StatusSeeOther)
		} else {
			fmt.Printf("Error fetching user: %v", err)
		}
	}
}
func (l *LikeHandler) DisLikeComment(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"LoggedIn": true,
	}
	var userID uuid.UUID
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
	
	postID := pathParts[2]
	commentID := pathParts[3]
	cookie, err := r.Cookie("session_id")
	if err != nil {
		data["LoggedIn"] = false
	}

	if err == nil && cookie != nil {
		sessionId := cookie.Value
		user, err := l.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			userID = user.ID
			l.LikeService.Create(userID,postID,commentID,"dislike")
			http.Redirect(w, r, fmt.Sprintf("http://localhost:8080/detailsPost/%v", postID), http.StatusSeeOther)
		} else {
			fmt.Printf("Error fetching user: %v", err)
		}
	}
}
