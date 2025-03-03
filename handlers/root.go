package handlers

import (
	"net/http"
	"text/template"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseGlob("./static/templates/*")
	tmp.Execute(w, "zero the goat!!!!")
}
func Login(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseGlob("./static/templates/*")
	tmp.ExecuteTemplate(w,"login.html",nil)
}
func Regist(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseGlob("./static/templates/*")
	tmp.ExecuteTemplate(w,"regist.html",nil)
}
