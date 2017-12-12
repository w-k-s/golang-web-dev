package main

import(
	"log"
	"net/http"
	_"html/template"
)

func main() {
	http.HandleFunc("/",index)
	http.Handle("/resources/",http.StripPrefix("/resources",http.FileServer(http.Dir("starting-files/public"))))
	log.Fatalln(http.ListenAndServe(":8080",nil))
}

func index(w http.ResponseWriter, req * http.Request){
	http.ServeFile(w,req,"starting-files/index.html")
}