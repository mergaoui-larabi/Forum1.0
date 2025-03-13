package handlers

import (
	"forum/config"
	"forum/database"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("session_token")
	if err != nil || sessionCookie.Value == "" {
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", map[string]bool{"Authenticated": false})
		return
	}

	_, exist := database.GetUserBySession(sessionCookie.Value)

	if !exist {
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", map[string]bool{"Authenticated": false})
		return
	}

	config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", map[string]bool{"Authenticated": true})
}
