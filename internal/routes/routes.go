package routes

import (
	"net/http"

	"forum/internal/handlers"
	"forum/internal/middleware"
	"forum/internal/utils"
)

func SetupRoutes(mux *http.ServeMux, authHandler *handlers.AuthHandler, postHandler *handlers.PostHandler, likeHandler *handlers.LikeHandler, authMiddleware *middleware.AuthMiddleware) {
	mux.HandleFunc("/static/", utils.SetupStaticFilesHandlers)

	mux.HandleFunc("/logout", authHandler.LogoutHandle)
	mux.HandleFunc("/register", authHandler.RegisterHandle)
	mux.HandleFunc("/login", authHandler.LoginHandle)

	mux.HandleFunc("/", postHandler.Home)

	mux.HandleFunc("/Posts/", postHandler.Posts)

	mux.HandleFunc("/create", postHandler.PostCreation)

	mux.HandleFunc("/createPost", postHandler.PostSaver)

	mux.HandleFunc("/sendcomment", postHandler.CommentSaver)
	mux.HandleFunc("/detailsPost/", postHandler.DetailsPost)
	mux.HandleFunc("/comment", postHandler.CommentGetter)
	mux.HandleFunc("/like/", likeHandler.LikePost)
	mux.HandleFunc("/dislike/", likeHandler.DisLikePost)
	mux.HandleFunc("/likeComment/", likeHandler.LikeComment)
	mux.HandleFunc("/dislikeComment/", likeHandler.DisLikeComment)
	mux.HandleFunc("/filters", postHandler.PostFilter)
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := http.Dir("./uploads").Open(r.URL.Path)
		if err != nil {
			utils.Error(w, http.StatusNotFound)
			return
		}
		http.FileServer(http.Dir("./uploads")).ServeHTTP(w, r)
	})))
	mux.HandleFunc("/javascript", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Referer") == "" {
			utils.Error(w, http.StatusNotFound)
			return
		}
		utils.Error(w, 1)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler, pattern := mux.Handler(r)
		if pattern == "" || pattern == "/" && r.URL.Path != "/" {
			utils.Error(w, http.StatusNotFound)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
