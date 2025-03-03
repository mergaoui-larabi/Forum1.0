package handlers

import (
	"net/http"
	"text/template"
)

func ForumHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./static/templates/*")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {	
		IsLoggedIn bool
	}{
		IsLoggedIn:true,
	}
	if err := tmpl.ExecuteTemplate(w,"index.html", data); err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseGlob("./static/templates/*")
	tmp.ExecuteTemplate(w, "login.html", nil)
}
func Regist(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseGlob("./static/templates/*")
	tmp.ExecuteTemplate(w, "regist.html", nil)
}
