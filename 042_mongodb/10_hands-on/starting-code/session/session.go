package session

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/waqqas-abdulkareem/golangwebdev/042_mongodb/10_hands-on/starting-code/models"
	"net/http"
	"time"
)

var DbUsers = map[string]models.User{}
var DbSessions = map[string]models.Session{}
var DbSessionsCleaned = time.Now()

const Length int = 30

func GetUser(w http.ResponseWriter, req *http.Request) models.User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = Length
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u models.User
	if s, ok := DbSessions[c.Value]; ok {
		s.LastActivity = time.Now()
		DbSessions[c.Value] = s
		u = DbUsers[s.UserName]
	}
	return u
}

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := DbSessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		DbSessions[c.Value] = s
	}
	_, ok = DbUsers[s.UserName]
	// refresh session
	c.MaxAge = Length
	http.SetCookie(w, c)
	return ok
}

func CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	ShowSessions()              // for demonstration purposes
	for k, v := range DbSessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(DbSessions, k)
		}
	}
	DbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	ShowSessions()             // for demonstration purposes
}

// for demonstration purposes
func ShowSessions() {
	fmt.Println("********")
	for k, v := range DbSessions {
		fmt.Println(k, v.UserName)
	}
	fmt.Println("")
}