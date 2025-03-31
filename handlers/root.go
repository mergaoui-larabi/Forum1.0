package handlers

import (
	"forum/config"
	authdatabase "forum/database/authentification"
	forumerror "forum/errors"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("session_token")
	if err != nil || sessionCookie.Value == "" {
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", map[string]any{"Authenticated": false})
		return
	}

	user_id, exist, err := authdatabase.SelectUserSession(sessionCookie.Value)

	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}

	if !exist {
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", map[string]any{"Authenticated": false})
		return
	}

	user, err := authdatabase.GetUserInfo(user_id)

	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}

	config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", map[string]any{"Authenticated": true, "Username": user.Username})
}
