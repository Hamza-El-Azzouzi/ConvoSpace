package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	// "forum/internal/handlers"

	"forum/internal/middleware"
	"forum/internal/services"
	"forum/internal/utils"
)

type PostHandler struct {
	AuthService     *services.AuthService
	AuthMidlaware   *middleware.AuthMidlaware
	CategoryService *services.CategoryService
	PostService     *services.PostService
	CommentService  *services.CommentService
	AuthHandler     *AuthHandler
}

func (p *PostHandler) HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, 405)
	}
	posts, err := p.PostService.AllPosts()

	categories, errCat := p.CategoryService.GetAllCategories()
	if errCat != nil {
		fmt.Printf("error kayn f categories getter : %v\n", err)
		utils.Error(w, 500)
	}
	
	if err != nil {
		fmt.Printf("error kayn f service POSt all : %v", err)
		utils.Error(w, 500)
	}
	data := map[string]interface{}{
		"LoggedIn":   true,
		"posts":      posts,
		"categories": categories,
	}
	isLogged, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		data["LoggedIn"] = isLogged
		data["Username"] = usermid.Username
	} else {
		data["LoggedIn"] = isLogged
	}

	utils.OpenHtml("index.html", w, data)
}

func (p *PostHandler) PostCreation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, 405)
	}
	categories, err := p.CategoryService.GetAllCategories()
	if err != nil {
		fmt.Printf("error kayn f categories getter : %v\n", err)
		utils.Error(w, 500)
	}
	data := map[string]any{
		"LoggedIn":   false,
		"categories": categories,
	}
	isLogged, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		data["LoggedIn"] = isLogged
		data["Username"] = usermid.Username
	} else {
		data["LoggedIn"] = isLogged
	}

	utils.OpenHtml("ask_question.html", w, data)
}

func (p *PostHandler) PostSaver(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Error(w, 405)
	}
	data := map[string]any{
		"LoggedIn": false,
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("error kayn fl form : %v", err)
		utils.Error(w, 500)
	}
	title := r.Form.Get("title")
	categories := r.Form["category"]
	subject := r.Form.Get("textarea")
	contentWithBreaks := strings.ReplaceAll(subject, "\n", "<br>")
	if title == "" || contentWithBreaks == "" || len(categories) == 0 {
		utils.Error(w, 400)
		return
	}
	isLogged, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		data["LoggedIn"] = isLogged
		data["Username"] = usermid.Username
		fmt.Println(categories)
		err = p.PostService.PostSave(usermid.ID, title, contentWithBreaks, categories)
		if err != nil {
			fmt.Printf("error kayn ftsrad dyal post : %v\n ", err)
			utils.Error(w, 500)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		data["LoggedIn"] = isLogged
	}

	utils.OpenHtml("index.html", w, data)
}

func (p *PostHandler) DetailsPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, 405)
	}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		utils.Error(w, 404)
		return
	}
	postID := pathParts[2]
	posts, err := p.PostService.GetPost(postID)
	data := map[string]any{
		"LoggedIn": false,
		"posts":    posts,
	}
	if err != nil {
		fmt.Println(err)
		utils.Error(w, 404)
		return
	}

	isLogged, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)

	if isLogged {
		data["LoggedIn"] = isLogged
		data["Username"] = usermid.Username
	} else {
		data["LoggedIn"] = isLogged
	}

	utils.OpenHtml("post-deatils.html", w, data)
}

func (p *PostHandler) CommentSaver(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Error(w, 405)
	}
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		utils.Error(w, 404)
	}
	postID := pathParts[2]
	err := r.ParseForm()
	if err != nil {
		utils.Error(w, 500)
	}
	commetContent := r.Form.Get("textarea")
	commetContentWithNewLines := strings.ReplaceAll(commetContent, "\n", "<br>")

	if commetContentWithNewLines == "" {
		utils.Error(w, 400)
		return
	}

	isLogged, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)

	if isLogged {
		err = p.CommentService.SaveComment(usermid.ID, postID, commetContentWithNewLines)
		if err != nil {
			fmt.Printf("error kayn ftsrad dyal comment : %v\n ", err)
			utils.Error(w, 500)
		}
	}

	http.Redirect(w, r, fmt.Sprintf("/detailsPost/%v", postID), http.StatusOK)
}

func (p *PostHandler) PostFilter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, 405)
	}
	filterby := r.URL.Query().Get("filterby")

	categorie := r.URL.Query().Get("categories")

	_, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)
	if usermid != nil {
		posts, err := p.PostService.FilterPost(filterby, categorie, usermid.ID)
		if err != nil {
			fmt.Printf("error kayn f filter : %v\n ", err)
			utils.Error(w, 500)
		}
		fmt.Println(posts)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	} else {
		utils.Error(w, 403)
	}
}
