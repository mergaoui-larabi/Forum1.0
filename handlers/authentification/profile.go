package handlers

import (
	"forum/config"
	authdatabase "forum/database/authentification"
	forumerror "forum/errors"
	"net/http"
)

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(userIDKey).(int)
	user, err := authdatabase.GetUserInfo(user_id)
	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}
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
