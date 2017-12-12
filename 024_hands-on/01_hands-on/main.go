package main

import(
	"fmt"
	"net/http"
	"log"
	"html/template"
)

func main() {

	http.HandleFunc("/",foo)
	http.HandleFunc("/dog",dog)
	http.HandleFunc("/dog.jpg",dogPic)

	log.Fatal(http.ListenAndServe(":8080",nil))
}

func foo(w http.ResponseWriter, req * http.Request){
	fmt.Fprint(w,"foo ran")
}

func dog(w http.ResponseWriter, req * http.Request){
	tpl := template.Must(template.ParseFiles("dog.gohtml"))
	log.Fatal(tpl.Execute(w,nil))	
}

func dogPic(w http.ResponseWriter, req * http.Request){
	http.ServeFile(w,req,"dog.jpg")
}