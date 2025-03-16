package handlers

import (
	"context"
	"fmt"
	"forum/config"
	"forum/database"
	"forum/security"
	"net/http"
)

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(userIDKey).(int)
	user := database.GetUserInfo(user_id)
	config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "profile.html", user) // when u excute 2 template the get concatinated one in top of the other
}

func UpddateProfile(w http.ResponseWriter, r *http.Request) {
	var errMap map[string]any
	value := r.PathValue("value")
	if r.Context().Value(errorCase) != nil {
		errMap = r.Context().Value(errorCase).(map[string]any)
	}
	switch value {
	case "username":
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "update.html", map[string]any{"username": true, "Error": errMap["Error"], "Message": errMap["Message"]})
		return
	case "email":
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "update.html", map[string]any{"email": true, "Error": errMap["Error"], "Message": errMap["Message"]})
		return
	case "password":
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "update.html", map[string]any{"password": true, "Error": errMap["Error"], "Message": errMap["Message"]})
		return
	default:
		http.Error(w, "bad req", http.StatusBadRequest)
	}
}

func SaveChanges(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(userIDKey).(int)
	switch r.PathValue("value") {
	case "username":

		new_username := r.FormValue("username")
		password := r.FormValue("current")
		if !config.ValidUsername(new_username) {
			ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Please enter a valid username"})
			UpddateProfile(w, r.WithContext(ctx))
			return
		}

		hash := database.GetUserHashById(user_id)
		if !security.CheckPassword(password, hash) {
			ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Please enter a valid username"})
			UpddateProfile(w, r.WithContext(ctx))
			return
		}

		if database.DupplicatedUsername(new_username) {
			ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Username Alredy exists try again"})
			UpddateProfile(w, r.WithContext(ctx))
			return
		}

		database.UpdateUsernmae(user_id, new_username)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	case "email":
		new_email := r.FormValue("email")
		password := r.FormValue("current")
		if !config.ValidEmail(new_email) {
			ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Invalid email try again"})
			UpddateProfile(w, r.WithContext(ctx))
			return
		}
		hash := database.GetUserHashById(user_id)
		if !security.CheckPassword(password, hash) {
			ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Wrong password"})
			UpddateProfile(w, r.WithContext(ctx))
			return
		}

		if database.DupplicatedEmail(new_email) {
			ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Email Alredy exists try again"})
			UpddateProfile(w, r.WithContext(ctx))
			return
		}

		database.UpdateEmail(user_id, new_email)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	case "password":
		current := r.FormValue("current")
		new := r.FormValue("new")
		confirm := r.FormValue("confirm")
		hash := database.GetUserHashById(user_id)
		if current == new {
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
			fmt.Println(err)
		}
		database.UpdatePassword(user_id, new_hash)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	default:
		http.Error(w, "bad req", http.StatusBadRequest)
		return
	}
}

func ServeDelete(w http.ResponseWriter, r *http.Request) {
	var errMap map[string]any
	if r.Context().Value(errorCase) != nil {
		errMap = r.Context().Value(errorCase).(map[string]any)
	}
	config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "delete.html", errMap)
}

func DeleteConfirmation(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(userIDKey).(int)
	password := r.FormValue("password")
	hash := database.GetUserHashById(user_id)
	if !security.CheckPassword(password, hash) {
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Wrong password"})
		ServeDelete(w, r.WithContext(ctx))
		return
	}
	LogoutHandler(w, r)
	database.DeleteUser(user_id)
}
