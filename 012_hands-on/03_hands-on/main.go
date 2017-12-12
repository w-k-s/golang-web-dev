package main

import(
	"log"
	"os"
	"text/template"
)

type Region string

const (
	Southern Region = "Southern"
	Central Region = "Central"
	Nothern Region = "Nothern"
)

type hotel struct{
	Name string
	Address string
	City string
	Zip string
	Region Region
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	
	hotels := []hotel{
		hotel{
			"Hotel Blue",
			"Blue Road",
			"Californian City #1",
			"10000",
			Southern,
		},
		hotel{
			"Hotel Red",
			"Red Road",
			"Californian City #2",
			"20000",
			Central,
		},
		hotel{
			"The Overlook Hotel",
			"Somwhere",
			"Sidewindder",
			"30000",
			Nothern,
		},
	}

	err := tpl.Execute(os.Stdout,hotels)
	if err != nil{
		log.Fatalln(err)
	}
}