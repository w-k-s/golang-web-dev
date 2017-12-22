package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/waqqas-abdulkareem/golangwebdev/042_mongodb/06_hands-on/starting-code/models"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"net/http"
	"os"
)

type UserController struct {
	db map[string]*models.User
}

func NewUserController(db map[string]*models.User) *UserController {
	uc := &UserController{db}
	uc.readDB()
	return uc
}

func (uc UserController) writeDB(){

	f, err := os.Create("db.json")
	if err != nil{
		panic(err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(&uc.db)
	if err != nil{
		panic(err)
	}
}

func (uc UserController) readDB(){
	f, err := os.Open("db.json")
	if err != nil{
		return
	}
	err = json.NewDecoder(f).Decode(&uc.db)
	if err != nil{
		panic(err)
	}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Fetch user
	u,ok := uc.db[id]
	if !ok {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	u.Id = uuid.NewV4().String()

	// store the user in mongodb
	uc.db[u.Id] = &u;

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)

	uc.writeDB()
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Delete user
	delete(uc.db,id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")

	uc.writeDB()
}
