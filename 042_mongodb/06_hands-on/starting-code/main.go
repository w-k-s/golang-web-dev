package main

import (
	"github.com/waqqas-abdulkareem/golangwebdev/042_mongodb/06_hands-on/starting-code/controllers"
	"github.com/waqqas-abdulkareem/golangwebdev/042_mongodb/06_hands-on/starting-code/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()
	db := map[string]*models.User{}
	// Get a UserController instance
	uc := controllers.NewUserController(db)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
