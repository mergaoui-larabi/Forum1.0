package handlers

import (
	"context"
	"forum/config"
	authdatabase "forum/database/authentification"
	forumerror "forum/errors"
	"forum/security"
	"net/http"
)

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
	hash, err := authdatabase.GetUserHashById(user_id)

	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}
	if !security.CheckPassword(password, hash) {
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Wrong password"})
		ServeDelete(w, r.WithContext(ctx))
		return
	}
	LogoutHandler(w, r)
	err = authdatabase.DeleteUser(user_id)

	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}
}
