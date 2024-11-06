package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	// "forum/internal/handlers"

	"forum/internal/services"
	"forum/internal/utils"
)

type PostHandler struct {
	AuthService     *services.AuthService
	CategoryService *services.CategoryService
	PostService     *services.PostService
	CommentService  *services.CommentService
}

func (p *PostHandler) HomeHandle(w http.ResponseWriter, r *http.Request) {
	posts, err := p.PostService.AllPosts()
	categories, errCat := p.CategoryService.GetAllCategories()
	if errCat != nil {
		fmt.Printf("error kayn f categories getter : %v\n", err)
		utils.Error(w,500)
	}
	// fmt.Println(posts)
	if err != nil {
		fmt.Printf("error kayn f service POSt all : %v", err)
		utils.Error(w,500)
	}
	data := map[string]interface{}{
		"LoggedIn":   true,
		"posts":      posts,
		"categories": categories,
	}
	cookie, err := r.Cookie("session_id")
	if err != nil {
		data["LoggedIn"] = false
	}

	if cookie != nil {
		sessionId := cookie.Value
		user, err := p.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {

			data["LoggedIn"] = true
			data["Username"] = user.Username

		} else {
			fmt.Printf("Error fetching user: %v", err)
			utils.Error(w,500)
		}
	}

	utils.OpenHtml("index.html", w, data)
}

func (p *PostHandler) PostCreation(w http.ResponseWriter, r *http.Request) {
	categories, err := p.CategoryService.GetAllCategories()
	if err != nil {
		fmt.Printf("error kayn f categories getter : %v\n", err)
		utils.Error(w,500)
	}
	data := map[string]any{
		"LoggedIn":   false,
		"categories": categories,
	}

	// Check if session_id cookie exists
	cookie, err := r.Cookie("session_id")
	if err == nil && cookie != nil {
		sessionId := cookie.Value

		user, err := p.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			data["LoggedIn"] = true
			data["Username"] = user.Username

		} else {
			fmt.Printf("Error fetching user: %v", err)
			utils.Error(w,500)
		}
	}
	utils.OpenHtml("ask_question.html", w, data)
}

func (p *PostHandler) PostSaver(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"LoggedIn": false,
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("error kayn fl form : %v", err)
		utils.Error(w,500)
	}
	title := r.Form.Get("title")
	categories := r.Form["category"]
	subject := r.Form.Get("textarea")
	contentWithBreaks := strings.ReplaceAll(subject, "\n", "<br>")
	cookie, err := r.Cookie("session_id")
	if err == nil && cookie != nil {
		sessionId := cookie.Value

		user, err := p.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			data["LoggedIn"] = true
			data["Username"] = user.Username

		} else {
			fmt.Printf("Error fetching user: %v", err)
			utils.Error(w,500)
		}
		err = p.PostService.PostSave(user.ID, title, contentWithBreaks, categories)
		if err != nil {
			fmt.Printf("error kayn ftsrad dyal post : %v\n ", err)
			utils.Error(w,500)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}

	utils.OpenHtml("index.html", w, data)
}

func (p *PostHandler) DetailsPost(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		utils.Error(w,404)
		return
	}
	postID := pathParts[2]
	posts, err := p.PostService.GetPost(postID)
	data := map[string]any{
		"LoggedIn": false,
		"posts":    posts,
	}
	if err != nil || posts.Username == "" {
		fmt.Println(err)
		utils.Error(w,404)
		return
	}
	cookie, err := r.Cookie("session_id")
	if err == nil && cookie != nil {
		sessionId := cookie.Value

		user, err := p.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			data["LoggedIn"] = true
			data["Username"] = user.Username

		} else {
			fmt.Printf("Error fetching user: %v", err)
			utils.Error(w,500)
		}
	}

	utils.OpenHtml("post-deatils.html", w, data)
}

func (p *PostHandler) CommentSaver(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		utils.Error(w,404)
	}
	postID := pathParts[2]
	err := r.ParseForm()
	if err != nil {
		utils.Error(w,500)
	}
	commetContent := r.Form.Get("textarea")
	commetContentWithNewLines := strings.ReplaceAll(commetContent, "\n", "<br>")
	cookie, err := r.Cookie("session_id")
	if err == nil && cookie != nil {
		sessionId := cookie.Value

		user, err := p.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err != nil {
			fmt.Printf("error fetch user : %v", err)
			utils.Error(w,500)
		}
		if len(commetContentWithNewLines) != 5 {
			err = p.CommentService.SaveComment(user.ID, postID, commetContentWithNewLines)
			if err != nil {
				fmt.Printf("error kayn ftsrad dyal comment : %v\n ", err)
				utils.Error(w,500)
			}
		}

	}

	http.Redirect(w, r, fmt.Sprintf("/detailsPost/%v", postID), http.StatusSeeOther)
}

func (p *PostHandler) PostFilter(w http.ResponseWriter, r *http.Request) {
	post := r.URL.Query().Get("posts")
	like := r.URL.Query().Get("likes")
	categories := strings.Split(r.URL.Query().Get("categories"), ",")
	fmt.Println(post)
	fmt.Println(like)
	fmt.Println(categories)
	if categories[0] == "" {
		categories = []string{}
	}
	posts, err := p.PostService.FilterPost(post, like, categories)
	if err != nil {
		fmt.Printf("error kayn f filter : %v\n ", err)
		utils.Error(w,500)
	}
	fmt.Println(posts)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
