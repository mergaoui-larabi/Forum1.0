package handlers

import (
	"context"
	"forum/config"
	authdatabase "forum/database/authentification"
	forumerror "forum/errors"
	"forum/security"
	"net/http"
	"time"
)

func SwitchLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ServLogin(w, r)
	case http.MethodPost:
		SubmitLogin(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func ServLogin(w http.ResponseWriter, r *http.Request) {
	var errMap map[string]any
	if r.Context().Value(errorCase) != nil {
		errMap = r.Context().Value(errorCase).(map[string]any)
	}
	config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "login.html", errMap)
}

func SubmitLogin(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	exist, err := authdatabase.AlreadyExists(username, username)

	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}

	if (!config.ValidUsername(username) && !config.ValidEmail(username)) || !config.ValidPassword(password) || !exist { //TODO: it should be a better way
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "invalid credentials try again"})
		ServLogin(w, r.WithContext(ctx))
		return
	}

	user_id, hash, err := authdatabase.GetUserHashByUsername(username)

	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}

	if !security.CheckPassword(password, hash) {
		ctx := context.WithValue(r.Context(), errorCase, map[string]any{"Error": true, "Message": "Wrong password try again"})
		ServLogin(w, r.WithContext(ctx))
		return
	}

	session := security.GenerateToken(32) // TODO: UUID bonus csrf implementation genrate csrf read it in front end js and match it with server go

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session,
		Expires:  time.Now().Add(time.Hour * 1),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	err = authdatabase.UpdateUserSession(user_id, session)

	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
