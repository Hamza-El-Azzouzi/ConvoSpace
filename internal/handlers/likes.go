package handlers

import (
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
	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed)
	}
	referer := r.Header.Get("Referer")

	if referer == "" {
		referer = "/"
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
		http.Redirect(w, r, referer, http.StatusSeeOther)
	} else {
		utils.Error(w, http.StatusForbidden)
	}
}

func (l *LikeHandler) DisLikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed)
	}
	referer := r.Header.Get("Referer")

	if referer == "" {
		referer = "/"
	}
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		http.Error(w, "Invalid URL", http.StatusNotFound)
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
		http.Redirect(w, r, referer, http.StatusSeeOther)
	} else {
		utils.Error(w, http.StatusForbidden)
	}
}

func (l *LikeHandler) LikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed)
	}
	referer := r.Header.Get("Referer")

	if referer == "" {
		referer = "/"
	}
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		http.Error(w, "Invalid URL", http.StatusNotFound)
	}

	commentID := pathParts[2]

	isLogged, usermid := l.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		err := l.LikeService.Create(usermid.ID, "", commentID, "like", true)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, referer, http.StatusSeeOther)
	} else {
		utils.Error(w, http.StatusForbidden)
	}
}

func (l *LikeHandler) DisLikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed)
	}
	referer := r.Header.Get("Referer")

	if referer == "" {
		referer = "/"
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
		http.Redirect(w, r, referer, http.StatusSeeOther)
	} else {
		utils.Error(w, http.StatusForbidden)
	}
}
