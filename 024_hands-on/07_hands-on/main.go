package main

import(
	"log"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/",index)
	http.Handle("/resources/",http.StripPrefix("/resources",http.FileServer(http.Dir("starting-files/public"))))
	log.Fatalln(http.ListenAndServe(":8080",nil))
}

func index(w http.ResponseWriter, req * http.Request){
	tpl := template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
	err := tpl.Execute(w,nil)
	if err != nil{
		log.Fatalln(err)
	}
}