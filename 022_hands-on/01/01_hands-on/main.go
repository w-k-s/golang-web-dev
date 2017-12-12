package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

func index(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w,"Index")
}

func dog(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w,"Dog")
}

func me(w http.ResponseWriter, req *http.Request){

	data := struct{
		Name string
		Mood string
	}{
		"W.K.S",
		"Depressed",
	}

	err := tpl.Execute(w,data)
	if err != nil{
		log.Fatalln(err)
	}
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("profile.gohtml"))
}

func main() {
	
	http.Handle("/",http.HandlerFunc(index))
	http.Handle("/dog/",http.HandlerFunc(dog))
	http.Handle("/me/",http.HandlerFunc(me))

	http.ListenAndServe(":8080",nil)
}