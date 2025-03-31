package handlers

import (
	authdatabase "forum/database/authentification"
	forumerror "forum/errors"
	"net/http"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	sessionCookie, err := r.Cookie("session_token")
	if err != nil || sessionCookie.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	hasSession, err := authdatabase.ResetUserSession(sessionCookie.Value)

	if err != nil {
		forumerror.TempErr(w, err, http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
	})

	if !hasSession {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
