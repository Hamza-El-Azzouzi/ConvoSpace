package handlers

import (
	"fmt"
	"net/http"
	"time"

	"forum/internal/services"
	"forum/internal/utils"

	"github.com/gofrs/uuid/v5"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func (h *AuthHandler) LogoutHandle(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		// If no session cookie exists, just redirect to home
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	sessionID := cookie.Value

	// Delete the session from the database
	_, err = h.AuthService.UserRepo.DB.Exec("DELETE FROM sessions WHERE session_id = ?", sessionID)
	if err != nil {
		http.Error(w, "Unable to log out", http.StatusInternalServerError)
		return
	}

	// Remove the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-1 * time.Hour), // Set expiration to past to delete the cookie
		HttpOnly: true,
	})

	// Redirect to the homepage or login page after logging out
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *AuthHandler) LoginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.OpenHtml("login.html", w, r)
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		email := r.Form.Get("email")
		password := r.Form.Get("password")

		// Check if passwords match

		// Check if fields are filled
		if email == "" || password == "" {
			fmt.Fprintln(w, "Please fill in all fields")
			return
		}

		// Process the user registration here (e.g., save to the database)

		user, loginError := h.AuthService.Login(email, password)
		if user == nil {
			fmt.Fprintln(w, "Invalid email or password")
			return
		}
		sessionID := uuid.Must(uuid.NewV4()).String()
		_, err = h.AuthService.UserRepo.DB.Exec("INSERT INTO sessions (session_id, user_id) VALUES (?, ?)", sessionID, user.ID)
		if err != nil {
			http.Error(w, "Unable to create session", http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})

		fmt.Println(user)
		if loginError != nil {
			fmt.Println(loginError)
			http.Error(w, "Unable to log in", http.StatusInternalServerError)
			return
		}
		fmt.Println("User loged successfully!")

		// Redirect to home page after successful registration
		http.Redirect(w, r, "/", http.StatusSeeOther) // Use StatusSeeOther for redirect after POST
	} else {
		utils.OpenHtml("login.html", w, r)
	}
}

func (h *AuthHandler) RegisterHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.OpenHtml("signup.html", w, r)
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		userName := r.Form.Get("Username")
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		confirmPassword := r.Form.Get("Confirm-Password")

		// Check if passwords match
		if password != confirmPassword {
			fmt.Fprintln(w, "Passwords do not match!")
			return
		}

		// Check if fields are filled
		if userName == "" || email == "" || password == "" {
			fmt.Fprintln(w, "Please fill in all fields")
			return
		}

		// Process the user registration here (e.g., save to the database)

		registrError := h.AuthService.Register(userName, email, password)
		if registrError != nil {
			fmt.Println(registrError)
			http.Error(w, "Unable to register user", http.StatusInternalServerError)
			return
		}
		fmt.Println("User registered successfully!")

		// Redirect to home page after successful registration
		http.Redirect(w, r, "/", http.StatusSeeOther) // Use StatusSeeOther for redirect after POST
	}
}

// func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
// 	var user models.User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	err = h.AuthService.Register(user.Username, user.Email, user.PasswordHash)
// 	if err != nil {
// 		http.Error(w, "Error registering user", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
// }

// func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
// 	var credentials struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&credentials)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	// token, err := h.Login(credentials.Username, credentials.Password)

// 	// if err != nil {
// 	// 	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 	// 	return
// 	// }

// 	w.WriteHeader(http.StatusOK)
// 	// json.NewEncoder(w).Encode(map[string]string{"token": token})
// }
