package handlers

import (
	"fmt"
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
		utils.Error(w, 405)
	}
	referer := r.Header.Get("Referer")

    if referer == "" {
        referer = "/"
    }
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		utils.Error(w, 404)
	}
	postID := pathParts[2]

	isLogged, usermid := l.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		l.LikeService.Create(usermid.ID, postID, "", "like", false)
		http.Redirect(w, r, referer, http.StatusSeeOther)
		return
	}else{
		utils.Error(w,403)
	}
}

func (l *LikeHandler) DisLikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, 405)
	}
	referer := r.Header.Get("Referer")

    if referer == "" {
        referer = "/"
    }
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
	postID := pathParts[2]

	isLogged, usermid := l.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		l.LikeService.Create(usermid.ID, postID, "", "dislike", false)
		http.Redirect(w, r, referer, http.StatusSeeOther)
		return

	}
}

func (l *LikeHandler) LikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, 405)
	}
	referer := r.Header.Get("Referer")

    if referer == "" {
        referer = "/"
    }
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
	fmt.Println(pathParts)
	commentID := pathParts[2]

	isLogged, usermid := l.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		err := l.LikeService.Create(usermid.ID, "", commentID, "like", true)
		if err != nil {
			fmt.Printf("err kayn f like comment : %v", err)
			utils.Error(w, 500)
		}
		http.Redirect(w, r, referer, http.StatusSeeOther)

	}
}

func (l *LikeHandler) DisLikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, 405)
	}
	referer := r.Header.Get("Referer")

    if referer == "" {
        referer = "/"
    }
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		utils.Error(w, 404)
	}

	commentID := pathParts[2]


	isLogged, usermid := l.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		err := l.LikeService.Create(usermid.ID, "", commentID, "dislike", true)
			if err != nil {
				fmt.Printf("err kayn f like comment : %v", err)
				utils.Error(w, 500)
			}
		http.Redirect(w, r, referer, http.StatusSeeOther)

	}
}
