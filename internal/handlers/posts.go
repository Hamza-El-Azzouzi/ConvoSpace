package handlers

import (
	"fmt"
	"net/http"

	// "forum/internal/handlers"

	"forum/internal/services"
	"forum/internal/utils"
)

type PostHandler struct {
	AuthService     *services.AuthService
	CategoryService *services.CategoryService
	PostService 	*services.PostService
}

func (p *PostHandler)HomeHandle(w http.ResponseWriter, r *http.Request) {
	posts ,err:= p.PostService.AllPosts()
	fmt.Println(posts)
	if err != nil{
		fmt.Printf("error kayn f service POSt all : %v", err)
	}
	data := map[string]interface{}{
		"LoggedIn": true,
		"posts":posts,
	}
	cookie, err := r.Cookie("session_id")
	if err != nil {
		data = map[string]interface{}{
			"LoggedIn": false,
			"posts":posts,
		}
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
		"LoggedIn":   false,
	}
	err := r.ParseForm()
	if err != nil{
		fmt.Printf("error kayn fl form : %v",err)
	}
	title := r.Form.Get("title")
	categories := r.Form["category"]
	subject := r.Form.Get("textarea")
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
		err = p.PostService.PostSave(user.ID,title,subject,categories) 
		if err != nil{
			fmt.Printf("error kayn ftsrad dyal post : %v\n ",err)
		}else{
			http.Redirect(w,r,"/",http.StatusSeeOther)
		}
		
	}

	// Render the template with user data
	utils.OpenHtml("index.html", w, data)
}