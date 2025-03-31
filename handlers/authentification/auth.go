package handlers

import (
	"context"
	"net/http"
	authdatabase "forum/database/authentification"
	forumerror "forum/errors"
)

type contextKey string

const userIDKey contextKey = "user_id"
const errorCase contextKey = "error_case"

func AuthMidleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("session_token")
		if err != nil || sessionCookie.Value == "" {
			ServLogin(w, r)
			return
		}

		user_id, exist, err := authdatabase.SelectUserSession(sessionCookie.Value)

		if err != nil {
			forumerror.TempErr(w, err, http.StatusInternalServerError)
		}

		if !exist {
			ServLogin(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, user_id) //avoid collisions
		next(w, r.WithContext(ctx))
	}
}
