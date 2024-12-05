package handlers

import (
	"net/http"
	"strings"

	"forum/internal/middleware"
	"forum/internal/services"
	"forum/internal/utils"
)

// import (
// 	"encoding/json"
// 	"net/http"

// 	"forum/internal/middleware"
// 	"forum/internal/services"
// )

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
	postid := r.URL.Path
	POSTid := strings.Split(postid, "/")
	if len(POSTid) != 3 {
		utils.Error(w, http.StatusNotFound)
	}
	ID := POSTid[2]

	_, user := l.AuthMidlaware.IsUserLoggedIn(w, r)

	if liked == "post" {
		err := l.LikeService.Create(user.ID, ID, "", typeOf, false)

		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
	} else {
		err := l.LikeService.Create(user.ID, "", ID, typeOf, true)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
	}

	// data, err := l.LikeService.GetLikes(ID, liked)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(data)

}
