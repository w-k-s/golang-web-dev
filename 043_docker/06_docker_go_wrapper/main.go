package main

import (
	"github.com/w-k-s/golangwebdev/043_docker/06_docker_go_wrapper/home"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
	"flag"
	"log"
	"fmt"
)

var port int
var dbConnString string
var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("./templates/*"))

	flag.IntVar(&port, "port", 8080, "Specify the port to listen to.")

	flag.Parse()

	log.Printf("Port: %d", port)
	log.Println("Init Complete")
}

func main() {
	r := mux.NewRouter()
	c := home.NewController(tpl)
	
	r.HandleFunc("/", c.Index).
		Methods("GET")

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

//3. Dependencies