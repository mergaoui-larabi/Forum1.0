package handlers

import (
	"forum/config"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", nil)
}
