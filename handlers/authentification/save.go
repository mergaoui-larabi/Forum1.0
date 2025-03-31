package handlers

import (
	"context"
	"forum/config"
	authdatabase "forum/database/authentification"
	forumerror "forum/errors"
	"forum/security"
	"net/http"
)

func SaveChanges(w http.ResponseWriter, r *http.Request) {
	switch r.PathValue("value") {
	case "username":
		SaveUsername(w, r)
		return
	case "email":
		SaveEmail(w, r)
		return
	case "password":
		SavePassword(w, r)
		return
	default:
		http.Error(w, "bad req", http.StatusBadRequest)
		return
	}
}

func SaveUsername(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(userIDKey).(int)
	new_username := r.FormValue("username")
	password := r.FormValue("current")
	if !config.ValidUsername(new_username) {
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Please enter a valid username"})
		UpddateProfile(w, r.WithContext(ctx))
		return
	}

	hash, err := authdatabase.GetUserHashById(user_id)
	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}
	if !security.CheckPassword(password, hash) {
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Wrong password"})
		UpddateProfile(w, r.WithContext(ctx))
		return
	}

	dupp, err := authdatabase.DupplicatedUsername(new_username)
	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}
	if dupp {
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Username Alredy exists try again"})
		UpddateProfile(w, r.WithContext(ctx))
		return
	}

	err = authdatabase.UpdateUsernmae(user_id, new_username)
	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func SaveEmail(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(userIDKey).(int)
	new_email := r.FormValue("email")
	password := r.FormValue("current")
	if !config.ValidEmail(new_email) {
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Invalid email try again"})
		UpddateProfile(w, r.WithContext(ctx))
		return
	}
	hash, err := authdatabase.GetUserHashById(user_id)
	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}
	if !security.CheckPassword(password, hash) {
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Wrong password"})
		UpddateProfile(w, r.WithContext(ctx))
		return
	}

	dupp, err := authdatabase.DupplicatedEmail(new_email)
	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}
	if dupp {
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Email Alredy exists try again"})
		UpddateProfile(w, r.WithContext(ctx))
		return
	}

	err = authdatabase.UpdateEmail(user_id, new_email)
	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func SavePassword(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(userIDKey).(int)
	current := r.FormValue("current")
	new := r.FormValue("new")
	confirm := r.FormValue("confirm")
	hash, err := authdatabase.GetUserHashById(user_id)

	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}
	if current == new || !config.ValidPassword(new) {
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "You used an Old password"})
		UpddateProfile(w, r.WithContext(ctx))
		return
	}
	if !security.CheckPassword(current, hash) {
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Wrong password"})
		UpddateProfile(w, r.WithContext(ctx))
		return
	}
	if new != confirm {
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Please Confirm Your password"})
		UpddateProfile(w, r.WithContext(ctx))
		return
	}
	new_hash, err := security.HashPassword(new)
	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}

	err = authdatabase.UpdatePassword(user_id, new_hash)
	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
