package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		utils.Error(w, http.StatusNotFound)
		return
	}
	pagination := pathParts[2]
	if pagination == "" {
		utils.Error(w, http.StatusNotFound)
		return
	}
	nPagination, err := strconv.Atoi(pagination)
	if err != nil {
		utils.Error(w, http.StatusNotFound)
		return
	}
	posts, err := p.PostService.AllPosts(nPagination)
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

	if title == "" || subject == "" || len(categories) == 0 {
		utils.Error(w, http.StatusBadRequest)
		return
	}
	isLogged, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)
	if isLogged {
		data["LoggedIn"] = isLogged
		data["Username"] = usermid.Username
		err = p.PostService.PostSave(usermid.ID, title, subject, categories)
		if err != nil {
			utils.Error(w, http.StatusBadRequest)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		data["LoggedIn"] = isLogged
	}
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
	if postID == "" {
		utils.Error(w, http.StatusNotFound)
		return
	}
	posts, err := p.PostService.GetPost(postID)
	if err != nil || posts.PostID == uuid.Nil {
		utils.Error(w, http.StatusNotFound)
		return
	}
	comment, err := p.CommentService.GetCommentByPost(postID, 0)
	if err != nil {
		utils.Error(w, http.StatusNotFound)
		return
	}
	data := map[string]any{
		"LoggedIn": false,
		"posts":    posts,
		"comment":  comment,
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
	var commentData models.CommentData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&commentData)
	
	defer r.Body.Close()

	if err != nil {
		utils.Error(w, http.StatusBadRequest)
		return
	}
	isLogged, userId := p.AuthMidlaware.IsUserLoggedIn(w, r)
	if !isLogged {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = p.CommentService.SaveComment(userId.ID, commentData.PostId, commentData.Comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comment, err := p.CommentService.GetCommentByPost(commentData.PostId, 0)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
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
	pagination := r.URL.Query().Get("pagination")
	nPagination, err := strconv.Atoi(pagination)
	if err != nil {
		utils.Error(w, http.StatusBadRequest)
		return
	}
	isLogged, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)

	if usermid != nil {
		filterby = r.URL.Query().Get("filterby")
	}
	if filterby != "" {
		posts, err = p.PostService.FilterPost(filterby, categorie, usermid.ID, nPagination)
		if err != nil {
			fmt.Println(err.Error())
			utils.Error(w, http.StatusInternalServerError)
			return
		}

	} else {
		posts, err = p.PostService.FilterPost(filterby, categorie, uuid.Nil, nPagination)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError)
			return
		}

	}

	data := map[string]any{
		"LoggedIn": false,
		"posts":    posts,
	}

	if isLogged {
		data["LoggedIn"] = isLogged
	} else {
		data["LoggedIn"] = isLogged
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (p *PostHandler) CommentGetter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed)
		return
	}
	var err error
	postID := r.URL.Query().Get("postId")
	pagination := r.URL.Query().Get("offset")
	nPagination, err := strconv.Atoi(pagination)
	if err != nil {
		utils.Error(w, http.StatusBadRequest)
		return
	}
	comment, err := p.CommentService.GetCommentByPost(postID, nPagination)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
}
