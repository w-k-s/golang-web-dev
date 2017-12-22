package main

import (
	"github.com/waqqas-abdulkareem/golangwebdev/042_mongodb/10_hands-on/starting-code/controllers"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	uc := controllers.NewUserController(tpl)

	http.HandleFunc("/", uc.Index)
	http.HandleFunc("/bar", uc.Bar)
	http.HandleFunc("/signup", uc.Signup)
	http.HandleFunc("/login", uc.Login)
	http.HandleFunc("/logout", uc.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}