package routes

import (
	"net/http"

	"forum/internal/handlers"
	"forum/internal/utils"
)

func SetupRoutes(mux *http.ServeMux, authHandler *handlers.AuthHandler, postHandler *handlers.PostHandler, likeHandler *handlers.LikeHandler) {
	mux.HandleFunc("/static/", utils.SetupStaticFilesHandlers)
	
	mux.HandleFunc("/", postHandler.HomeHandle)
	mux.HandleFunc("/create", postHandler.PostCreation)
	mux.HandleFunc("/createPost", postHandler.PostSaver)
	mux.HandleFunc("/sendcomment/", postHandler.CommentSaver)
	mux.HandleFunc("/logout", authHandler.LogoutHandle)
	mux.HandleFunc("/login", authHandler.LoginHandle)
	mux.HandleFunc("/register", authHandler.RegisterHandle)
	mux.HandleFunc("/detailsPost/", postHandler.DetailsPost)
	mux.HandleFunc("/like/", likeHandler.LikePost)
	mux.HandleFunc("/dislike/", likeHandler.DisLikePost)
	mux.HandleFunc("/likeComment/", likeHandler.LikeComment)
	mux.HandleFunc("/dislikeComment/", likeHandler.DisLikeComment)
	mux.HandleFunc("/filters", postHandler.PostFilter)
	mux.HandleFunc("/checker", authHandler.CheckDoubleLogging)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler, pattern := mux.Handler(r)
		if pattern == "" || pattern == "/" && r.URL.Path != "/" {
			utils.Error(w, 404)
			return
		}
		handler.ServeHTTP(w, r)
	})
}