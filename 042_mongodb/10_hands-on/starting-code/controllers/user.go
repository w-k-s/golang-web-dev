package controllers

import(
	"html/template"
	"github.com/waqqas-abdulkareem/golangwebdev/042_mongodb/10_hands-on/starting-code/models"
	"github.com/waqqas-abdulkareem/golangwebdev/042_mongodb/10_hands-on/starting-code/session"
	"net/http"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserController struct{
	tpl *template.Template
}

func NewUserController(tpl *template.Template) *UserController{
	return &UserController{
		tpl,
	}
}

func (uc UserController) Index(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	session.ShowSessions() // for demonstration purposes
	uc.tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func (uc UserController) Bar(w http.ResponseWriter, req *http.Request) {
	
	u := session.GetUser(w, req)
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	session.ShowSessions() // for demonstration purposes
	uc.tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func (uc UserController) Signup(w http.ResponseWriter, req *http.Request) {

	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		if _, ok := session.DbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = session.Length
		http.SetCookie(w, c)
		session.DbSessions[c.Value] = models.Session{un, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = models.User{un, bs, f, l, r}
		session.DbUsers[un] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.ShowSessions() // for demonstration purposes
	uc.tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func (uc UserController) Login(w http.ResponseWriter, req *http.Request) {
	
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := session.DbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = session.Length
		http.SetCookie(w, c)
		session.DbSessions[c.Value] = models.Session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.ShowSessions() // for demonstration purposes
	uc.tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func (uc UserController) Logout(w http.ResponseWriter, req *http.Request) {
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(session.DbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(session.DbSessionsCleaned) > (time.Second * 30) {
		go session.CleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
