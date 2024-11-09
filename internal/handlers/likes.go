package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"forum/internal/services"
	"forum/internal/utils"

	"github.com/gofrs/uuid/v5"
)

type LikeHandler struct {
	LikeService *services.LikeService
	AuthService *services.AuthService
}

func (l *LikeHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet{
		utils.Error(w,405)
	}
	var userID uuid.UUID
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		utils.Error(w,404)
	}
	postID := pathParts[2]
	cookie, err := r.Cookie("session_id")

	if err == nil && cookie != nil {
		sessionId := cookie.Value
		user, err := l.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			userID = user.ID
			l.LikeService.Create(userID, postID, "", "like", false)
			http.Redirect(w, r, fmt.Sprintf("http://10.1.2.1:8080/detailsPost/%v", postID), http.StatusSeeOther)
		} else {
			fmt.Printf("Error fetching user: %v", err)
			utils.Error(w, 500)
		}
	}
}

func (l *LikeHandler) DisLikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet{
		utils.Error(w,405)
	}
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
			l.LikeService.Create(userID, postID, "", "dislike", false)
			http.Redirect(w, r, fmt.Sprintf("http://10.1.2.1:8080/detailsPost/%v", postID), http.StatusSeeOther)
		} else {
			fmt.Printf("Error fetching user: %v", err)
			utils.Error(w,500)
		}
	}
}

func (l *LikeHandler) LikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet{
		utils.Error(w,405)
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

	if err == nil && cookie != nil {
		sessionId := cookie.Value
		user, err := l.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			userID = user.ID
			err = l.LikeService.Create(userID, "", commentID, "like", true)
			if err != nil{
				fmt.Printf("err kayn f like comment : %v", err)
				utils.Error(w,500)
			}
			
			http.Redirect(w, r, fmt.Sprintf("http://10.1.2.1:8080/detailsPost/%v", postID), http.StatusSeeOther)
		} else {
			fmt.Printf("Error fetching user: %v", err)
			utils.Error(w,500)
		}
	}
}

func (l *LikeHandler) DisLikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet{
		utils.Error(w,405)
	}
	var userID uuid.UUID
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		utils.Error(w,404)
	}

	postID := pathParts[2]
	commentID := pathParts[3]
	cookie, err := r.Cookie("session_id")

	if err == nil && cookie != nil {
		sessionId := cookie.Value
		user, err := l.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			userID = user.ID
			err = l.LikeService.Create(userID, "", commentID, "dislike", true)
			if err != nil{
				fmt.Printf("err kayn f like comment : %v", err)
				utils.Error(w,500)
			}
			http.Redirect(w, r, fmt.Sprintf("http://10.1.2.1:8080/detailsPost/%v", postID), http.StatusSeeOther)
		} else {
			fmt.Printf("Error fetching user: %v", err)
			utils.Error(w,500)
		}
	}
}
