package main

import (
	"os"
	"log"
	"text/template"
)

type something struct{
	Date string
	High float64
	Low float64
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	//parse csv

	err := tpl.Execute(os.Stdout,nil)
	if err != nil{
		log.Fatalln(err)
	}
}