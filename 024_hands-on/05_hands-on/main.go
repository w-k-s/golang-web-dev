package main

import (
	"log"
	"net/http"
	"html/template"
)	

func main() {
	fs := http.FileServer(http.Dir("public"))

	http.HandleFunc("/",index)
	http.Handle("/pics/",fs)

	log.Fatalln(http.ListenAndServe(":8080",nil))
}

func index(w http.ResponseWriter, req * http.Request){
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}