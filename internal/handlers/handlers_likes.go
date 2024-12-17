package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"forum/internal/middleware"
	"forum/internal/services"
	"forum/internal/utils"
)

type LikeHandler struct {
	LikeService   *services.LikeService
	AuthMidlaware *middleware.AuthMiddleware
	mutex         sync.Mutex
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

func (l *LikeHandler) react(w http.ResponseWriter, r *http.Request, liked, typeOfReact string) {
	l.mutex.Lock()
	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}
	postid := r.URL.Path
	POSTid := strings.Split(postid, "/")
	if len(POSTid) != 3 {
		utils.Error(w, http.StatusNotFound)
		return
	}
	ID := POSTid[2]
	logeddUser, user := l.AuthMidlaware.IsUserLoggedIn(w, r)
	if logeddUser {
		if liked == "post" {
			err := l.LikeService.Create(user.ID, ID, "", typeOfReact, liked)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		} else {
			err := l.LikeService.Create(user.ID, "", ID, typeOfReact, liked)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		data, err := l.LikeService.GetLikes(ID, liked)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
	} else {
		utils.Error(w, http.StatusForbidden)
	}
	l.mutex.Unlock()
}
