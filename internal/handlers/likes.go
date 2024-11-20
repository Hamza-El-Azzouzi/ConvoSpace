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
	l.react(w, r, "post", "like")
}

func (l *LikeHandler) DisLikePost(w http.ResponseWriter, r *http.Request) {
	l.react(w, r, "post", "dislike")
}

func (l *LikeHandler) LikeComment(w http.ResponseWriter, r *http.Request) {
	l.react(w, r, "comment", "like")
}

func (l *LikeHandler) DisLikeComment(w http.ResponseWriter, r *http.Request) {
	l.react(w, r, "comment", "dislike")
}

func (l *LikeHandler) react(w http.ResponseWriter, r *http.Request, liked, typeOf string) {
	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		utils.Error(w, http.StatusNotFound)
		return
	}

	ID := pathParts[2]

	isLogged, usermid := l.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		var err error
		if liked == "post" {
			err = l.LikeService.Create(usermid.ID, ID, "", typeOf, false)
		} else {
			err = l.LikeService.Create(usermid.ID, "", ID, typeOf, true)
		}
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		data, err := l.LikeService.GetLikes(ID, liked)
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
