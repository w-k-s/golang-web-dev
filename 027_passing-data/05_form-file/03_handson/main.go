package main

import(
	"fmt"
	"net/http"
	"html/template"
	_"io"
	"io/ioutil"
	"os"
	"path/filepath"
	_"mime/multipart"
	_"strings"
)

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/register",register)
	http.HandleFunc("/profile",profile)
	http.Handle("/pictures/",http.StripPrefix("/pictures",http.FileServer(http.Dir("./pictures"))))
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, req *http.Request){
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	err := tpl.Execute(w,nil)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}

func register(w http.ResponseWriter, req *http.Request){
	
	name := req.FormValue("username")
	file,_,err := req.FormFile("profile_picture")
	
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	dest,err := os.Create(filepath.Join("./pictures/",name))
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	_, err = dest.Write(contents)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	redirect := fmt.Sprintf("/profile?name=%s",name)
	http.Redirect(w,req,redirect,http.StatusSeeOther)
}

func profile(w http.ResponseWriter, req * http.Request){
	name := req.FormValue("name")
	if len(name) == 0{
		http.Error(w,"Name parameter missing",http.StatusInternalServerError)
		return
	}

	data := struct{
		Name string
		Url string
	}{
		name,
		fmt.Sprintf("pictures/%s",name),
	}

	tpl := template.Must(template.ParseFiles("./templates/profile.gohtml"))
	err := tpl.Execute(w,data)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
}