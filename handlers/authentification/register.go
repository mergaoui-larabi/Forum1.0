package handlers

import (
	"context"
	"forum/config"
	authdatabase "forum/database/authentification"
	"forum/security"
	"net/http"
)

func SwitchRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ServRegister(w, r)
	case http.MethodPost:
		SubmitRegister(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func ServRegister(w http.ResponseWriter, r *http.Request) {
	var errMap map[string]any
	if r.Context().Value(errorCase) != nil {
		errMap = r.Context().Value(errorCase).(map[string]any)
	}
	config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "register.html", errMap)
}

func SubmitRegister(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirm_password := r.FormValue("confirm_password")

	if !config.ValidUsername(username) || !config.ValidEmail(email) || !config.ValidPassword(password) || confirm_password != password { //TODO: it should be a better way
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "invalid credentials try again"})
		ServRegister(w, r.WithContext(ctx))
		return
	}

	hash, err := security.HashPassword(password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = authdatabase.AddNewUser(username, email, hash)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" || err.Error() == "UNIQUE constraint failed: users.email" {
			ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "username or email alredy used"})
			ServRegister(w, r.WithContext(ctx))
			return
		} else {
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
		}
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
