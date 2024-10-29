package handlers

import (
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
	// fmt.Println(posts)
	if err != nil {
		fmt.Printf("error kayn f service POSt all : %v", err)
	}
	data := map[string]interface{}{
		"LoggedIn": true,
		"posts":    posts,
	}
	cookie, err := r.Cookie("session_id")
	if err != nil {
		data["LoggedIn"] = false
	}

	if cookie != nil {
		data["LoggedIn"] = true
	}
	// Retrieve user from database using userID
	utils.OpenHtml("index.html", w, data)
}

func (p *PostHandler) PostCreation(w http.ResponseWriter, r *http.Request) {
	categories, err := p.CategoryService.GetAllCategories()
	if err != nil {
		fmt.Printf("error kayn f categories getter : %v\n", err)
	}
	// if len(categories) != 0 {
	// fmt.Println(categories)
	// }

	data := map[string]any{
		"LoggedIn":   false,
		"categories": categories,
	}

	// Check if session_id cookie exists
	cookie, err := r.Cookie("session_id")
	if err == nil && cookie != nil {
		sessionId := cookie.Value

		// Attempt to get user from session ID

		user, err := p.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			data["LoggedIn"] = true
			data["Username"] = user.Username // Assign username to "Username" key

		} else {
			fmt.Printf("Error fetching user: %v", err)
		}
	}

	// Render the template with user data
	utils.OpenHtml("ask_question.html", w, data)
}

func (p *PostHandler) PostSaver(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"LoggedIn": false,
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("error kayn fl form : %v", err)
	}
	title := r.Form.Get("title")
	categories := r.Form["category"]
	subject := r.Form.Get("textarea")
	contentWithBreaks := strings.ReplaceAll(subject, "\n", "<br>")
	// Check if session_id cookie exists
	cookie, err := r.Cookie("session_id")
	if err == nil && cookie != nil {
		sessionId := cookie.Value

		// Attempt to get user from session ID

		user, err := p.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			data["LoggedIn"] = true
			data["Username"] = user.Username

		} else {
			fmt.Printf("Error fetching user: %v", err)
		}
		err = p.PostService.PostSave(user.ID, title, contentWithBreaks, categories)
		if err != nil {
			fmt.Printf("error kayn ftsrad dyal post : %v\n ", err)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}

	// Render the template with user data
	utils.OpenHtml("index.html", w, data)
}

func (p *PostHandler) DetailsPost(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
	postID := pathParts[2]
	posts, err := p.PostService.GetPost(postID)
	data := map[string]any{
		"LoggedIn": false,
		"posts":    posts,
	}
	if err != nil {
		fmt.Println(err)
	}
	cookie, err := r.Cookie("session_id")
	if err == nil && cookie != nil {
		sessionId := cookie.Value

		// Attempt to get user from session ID

		user, err := p.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err == nil && user != nil {
			data["LoggedIn"] = true
			data["Username"] = user.Username

		} else {
			fmt.Printf("Error fetching user: %v", err)
		}
	}

	fmt.Println(posts)

	utils.OpenHtml("post-deatils.html", w, data)
}

func (p *PostHandler) CommentSaver(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		http.Error(w, "Invalid URL", 404)
	}
	postID := pathParts[2]
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("error kayn fl form dyal comment : %v", err)
	}
	commetContent := r.Form.Get("textarea")
	commetContentWithNewLines := strings.ReplaceAll(commetContent, "\n", "<br>")
	cookie, err := r.Cookie("session_id")
	if err == nil && cookie != nil {
		sessionId := cookie.Value

		// Attempt to get user from session ID

		user, err := p.AuthService.UserRepo.GetUserBySessionID(sessionId)
		if err!= nil{
			fmt.Printf("error fetch user : %v", err)
		}
		if len(commetContentWithNewLines) != 5 {
			err = p.CommentService.SaveComment(user.ID, postID, commetContentWithNewLines)
			if err != nil {
				fmt.Printf("error kayn ftsrad dyal comment : %v\n ", err)
			}
		}

	}

	http.Redirect(w,r,fmt.Sprintf("/detailsPost/%v",postID),http.StatusSeeOther)
}