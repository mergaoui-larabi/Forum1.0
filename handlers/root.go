package handlers

import (
	"net/http"
	"text/template"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseGlob("./static/templates/*")
	tmp.Execute(w, "zero the goat!!!!")
}
