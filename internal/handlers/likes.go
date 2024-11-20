package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"forum/internal/middleware"
	"forum/internal/services"
	"forum/internal/utils"
)

type LikeHandler struct {
	LikeService   *services.LikeService
	AuthService   *services.AuthService
	AuthMidlaware *middleware.AuthMiddleware
}

func (l *LikeHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		utils.Error(w, http.StatusNotFound)
		return
	}
	postID := pathParts[2]

	isLogged, usermid := l.AuthMidlaware.IsUserLoggedIn(w, r)

	if isLogged {
		err := l.LikeService.Create(usermid.ID, postID, "", "like", false)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		data, err := l.LikeService.GetLikesPost(postID)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	} else {
		utils.Error(w, http.StatusForbidden)
	}
}

func (l *LikeHandler) DisLikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		utils.Error(w,http.StatusNotFound)
		return
	}
	postID := pathParts[2]

	isLogged, usermid := l.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		err := l.LikeService.Create(usermid.ID, postID, "", "dislike", false)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		data, err := l.LikeService.GetLikesPost(postID)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	} else {
		utils.Error(w, http.StatusForbidden)
	}
}

func (l *LikeHandler) LikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}
	
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		utils.Error(w,http.StatusNotFound)
		return
	}

	commentID := pathParts[2]

	isLogged, usermid := l.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		err := l.LikeService.Create(usermid.ID, "", commentID, "like", true)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		data, err := l.LikeService.GetLikesComment(commentID)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	} else {
		utils.Error(w, http.StatusForbidden)
	}
}

func (l *LikeHandler) DisLikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed)
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		utils.Error(w, http.StatusNotFound)
	}

	commentID := pathParts[2]

	isLogged, usermid := l.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		err := l.LikeService.Create(usermid.ID, "", commentID, "dislike", true)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		data, err := l.LikeService.GetLikesComment(commentID)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	} else {
		utils.Error(w, http.StatusForbidden)
	}
}
