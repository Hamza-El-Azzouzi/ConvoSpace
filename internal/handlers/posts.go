package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	// "forum/internal/handlers"

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

func (p *PostHandler) HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, 405)
	}
	posts, err := p.PostService.AllPosts()

	if err != nil {
		fmt.Printf("error kayn f service POSt all : %v", err)
		utils.Error(w, 500)
		return
	}

	categories, errCat := p.CategoryService.GetAllCategories()
	if errCat != nil {
		fmt.Printf("error kayn f categories getter : %v\n", err)
		utils.Error(w, 500)
		return
	}

	
	data := map[string]interface{}{
		"LoggedIn":   true,
		"categories": categories,
	}
	isLogged, usermid := p.AuthMidlaware.IsUserLoggedIn(w, r)

	var postsWithStatus []map[string]interface{}
    for _, post := range posts {
        postData := map[string]interface{}{
            "PostID":        post.PostID,
            "Title":         post.Title,
            "Content":       post.Content,
            "CreatedAt":     post.CreatedAt,
            "UserID":        post.UserID,
            "Username":      post.Username,
            "Email":         post.Email,
            "FormattedDate": post.FormattedDate,
            "CategoryName":  post.CategoryName,
            "CommentCount":  post.CommentCount,
            "LikeCount":     post.LikeCount,
            "DisLikeCount":  post.DisLikeCount,
            "LoggedInP":      isLogged, // Add dynamic LoggedIn property
        }
        postsWithStatus = append(postsWithStatus, postData)
    }
	if isLogged {
		data["LoggedIn"] = isLogged
		data["Username"] = usermid.Username
		data["posts"] = postsWithStatus
	} else {
		data["LoggedIn"] = isLogged
		data["posts"] = postsWithStatus
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

	http.Redirect(w, r, fmt.Sprintf("/detailsPost/%v", postID), http.StatusSeeOther)
}

func (p *PostHandler) PostFilter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, 405)
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
			fmt.Printf("error kayn f filter : %v\n ", err)
			utils.Error(w, 500)
		}
		
	} else {
		posts, err = p.PostService.FilterPost(filterby, categorie, uuid.Nil)
		if err != nil {
			fmt.Printf("error kayn f filter : %v\n ", err)
			utils.Error(w, 500)
		}
		
	}
	fmt.Println(posts)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
