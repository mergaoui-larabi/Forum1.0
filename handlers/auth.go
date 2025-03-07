package handlers

import (
	"fmt"
	"forum/config"
	"forum/security"
	"net/http"
	"time"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

var Users = map[string]Login{}

func FromHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login-form":
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "login.html", nil)
	case "/register-form":
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "register.html", nil)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(Users)

	username := r.FormValue("email_user")
	password := r.FormValue("password")

	fmt.Println("u:", username, "p:", password)

	user, ok := Users[username]
	fmt.Println(user)
	fmt.Println(ok)
	fmt.Println(security.CheckPassword(password, user.HashedPassword))

	if !ok || !security.CheckPassword(password, user.HashedPassword) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}

	sessiontoken := security.GenerateToken(32)
	csrftoken := security.GenerateToken(32)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessiontoken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrftoken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	// need data base implemntaion
	user.SessionToken = sessiontoken
	user.CSRFToken = csrftoken
	Users[username] = user

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) < 8 || len(password) < 8 {
		http.Error(w, http.StatusText(http.StatusNonAuthoritativeInfo), http.StatusNotAcceptable)
		return
	}

	if _, ok := Users[username]; ok {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}
	hashpass, _ := security.HashPassword(password)

	Users[username] = Login{
		HashedPassword: hashpass,
	}

	fmt.Fprint(w, "u did it")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "register.html", nil)
}

func AuthorizedHandler(w http.ResponseWriter, r *http.Request) {
	config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "register.html", nil)
}
