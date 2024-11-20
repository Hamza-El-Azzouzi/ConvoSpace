package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"forum/internal/middleware"
	"forum/internal/models"
	"forum/internal/services"
	"forum/internal/utils"

	"github.com/gofrs/uuid/v5"
)

type PostHandler struct {
	AuthService     *services.AuthService
	AuthMidlaware   *middleware.AuthMiddleware
	CategoryService *services.CategoryService
	PostService     *services.PostService
	CommentService  *services.CommentService
	AuthHandler     *AuthHandler
}

func (p *PostHandler) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}
	utils.OpenHtml("index.html", w, nil)
}

func (p *PostHandler) Posts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}
	posts, err := p.PostService.AllPosts()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError)
		return
	}

	categories, errCat := p.CategoryService.GetAllCategories()
	if errCat != nil {
		utils.Error(w, http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"LoggedIn":   true,
		"categories": categories,
		"posts":      posts,
	}
	isLogged, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)

	if isLogged {
		data["LoggedIn"] = isLogged
		data["Username"] = usermid.Username
	} else {
		data["LoggedIn"] = isLogged
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (p *PostHandler) PostCreation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}

	categories, err := p.CategoryService.GetAllCategories()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError)
		return
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
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}
	data := map[string]any{
		"LoggedIn": false,
	}
	err := r.ParseForm()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError)
		return
	}
	title := r.Form.Get("title")
	categories := r.Form["category"]
	subject := r.Form.Get("textarea")
	contentWithBreaks := strings.ReplaceAll(subject, "\n", "<br>")

	if title == "" || contentWithBreaks == "" || len(categories) == 0 {
		utils.Error(w, http.StatusBadRequest)
		return
	}
	isLogged, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		data["LoggedIn"] = isLogged
		data["Username"] = usermid.Username
		err = p.PostService.PostSave(usermid.ID, title, contentWithBreaks, categories)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
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
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		utils.Error(w, http.StatusNotFound)
		return
	}
	postID := pathParts[2]
	posts, err := p.PostService.GetPost(postID)
	if err != nil {
		utils.Error(w, http.StatusNotFound)
		return
	}
	data := map[string]any{
		"LoggedIn": false,
		"posts":    posts,
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
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}
	var commentData struct {
		Content string `json:"content"`
		PostID  string `json:"postID"`
	}

	err := json.NewDecoder(r.Body).Decode(&commentData)
	if err != nil {
		utils.Error(w, http.StatusBadRequest)
		return
	}

	commetContentWithNewLines := strings.ReplaceAll(commentData.Content, "\n", "<br>")
	isLogged, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)

	if isLogged {
		err = p.CommentService.SaveComment(usermid.ID, commentData.PostID, commetContentWithNewLines)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		comment, err := p.CommentService.GetCommentByPost(commentData.PostID)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(comment)

	}
}

func (p *PostHandler) PostFilter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}

	filterby := ""
	var posts []models.PostWithUser
	var err error
	categorie := r.URL.Query().Get("categories")

	_, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)
	if usermid != nil {
		filterby = r.URL.Query().Get("filterby")
	}
	if filterby != "" {
		posts, err = p.PostService.FilterPost(filterby, categorie, usermid.ID)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}

	} else {
		posts, err = p.PostService.FilterPost(filterby, categorie, uuid.Nil)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
