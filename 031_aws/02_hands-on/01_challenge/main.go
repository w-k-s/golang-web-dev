package main

import(
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"

)

const CookieNameSession = "Session"
const CookieMaxAgeSession = 30

type User struct{
	FirstName string
	LastName string
	Username string
	Password []byte
	IsAdmin bool
}

var tpl *template.Template
var users map[string]User
var sessions map[string]string

func init(){
	users = make(map[string]User)
	sessions = make(map[string]string)
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/login",login)
	http.HandleFunc("/logout",logout)
	http.HandleFunc("/register",register)
	http.HandleFunc("/admin",admin)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":80",nil)
}

func index(w http.ResponseWriter, req * http.Request){
	var user *User
	var loggedIn bool
	if user,loggedIn = isLoggedIn(w,req); !loggedIn{
		http.Redirect(w,req,"/login",http.StatusSeeOther)
		return
	}

	data := struct{
		User *User
	}{
		user,
	}

	err := tpl.ExecuteTemplate(w,"index.gohtml",data)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}

func login(w http.ResponseWriter, req * http.Request){
	if _,loggedIn := isLoggedIn(w,req); loggedIn{
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost{
		username := req.FormValue("username")
		password := req.FormValue("password")

		if len(username) == 0 || len(password) == 0{
			http.Error(w,"Invalid Form",http.StatusBadRequest)
			return
		}

		var user User
		var ok bool
		if user,ok = users[username]; !ok{
			http.Redirect(w,req,"/register",http.StatusSeeOther)
			return
		}

		err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err != nil{
			http.Error(w,"Incorrect Username or Password",http.StatusBadRequest)
			return
		}

		cookie := &http.Cookie{
			Name: "Session",
			Value: uuid.NewV4().String(),
			HttpOnly: true,
		}
		sessions[cookie.Value] = username

		http.SetCookie(w,cookie)
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	err := tpl.ExecuteTemplate(w,"login.gohtml",nil)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}

func logout(w http.ResponseWriter, req * http.Request){
	cookie, err := req.Cookie("Session")
	if err != nil{
		http.Redirect(w,req,"/login",http.StatusSeeOther)
		return
	}

	cookie.MaxAge = -1

	http.SetCookie(w,cookie)
	http.Redirect(w,req,"/login",http.StatusSeeOther)
}

func register(w http.ResponseWriter, req * http.Request){
	if _,loggedIn := isLoggedIn(w,req); loggedIn{
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost{
		username := req.FormValue("username")
		password := req.FormValue("password")
		firstName := req.FormValue("first_name")
		lastName := req.FormValue("last_name")
		admin := req.FormValue("admin") == "1"

		fmt.Printf("adm: %v\n",admin)

		//validate params
		if len(username) == 0 || len(password) == 0 || len(firstName) == 0 || len(lastName) == 0{
			fmt.Printf("fn:%s ln:%s un:%s pw:%s\n",firstName,lastName,username,password)
			http.Error(w, "Form Invalid",http.StatusBadRequest)
			return
		}

		//check username unique
		if _,ok := users[username]; ok{
			http.Error(w,"Username taken",http.StatusBadRequest)
			return
		}

		//hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.MinCost)
		if err != nil{
			http.Error(w,"Couldn't secure password",http.StatusInternalServerError)
			return
		}

		user := User{
			FirstName: firstName,
			LastName: lastName,
			Username: username,
			Password: hashedPassword,
			IsAdmin: admin,
		}

		cookie := &http.Cookie{
			Name: CookieNameSession,
			Value: uuid.NewV4().String(),
			HttpOnly: true,
			MaxAge: CookieMaxAgeSession,
		}
		users[username] = user
		sessions[cookie.Value] = username

		http.SetCookie(w,cookie)
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	err := tpl.ExecuteTemplate(w,"register.gohtml",nil)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}

func admin(w http.ResponseWriter, req * http.Request){
	var user *User
	var loggedIn bool
	if user,loggedIn = isLoggedIn(w,req); !loggedIn{
		http.Redirect(w,req,"/login",http.StatusSeeOther)
		return
	}

	if !user.IsAdmin{
		http.Error(w,"Not an admin",http.StatusForbidden)
		return
	}

	data := struct{
		User *User
	}{
		user,
	}

	err := tpl.ExecuteTemplate(w,"admin.gohtml",data)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}

func isLoggedIn(w http.ResponseWriter, req *http.Request) (*User,bool){
	cookie, err := req.Cookie(CookieNameSession)
	if err != nil{
		return nil,false
	}

	var username string
	var ok bool
	if username,ok = sessions[cookie.Value];!ok{
		return nil,false
	}

	var user User
	if user,ok = users[username]; !ok{
		return nil,false
	}

	cookie.MaxAge = CookieMaxAgeSession
	http.SetCookie(w,cookie)
	return &user,true
}

