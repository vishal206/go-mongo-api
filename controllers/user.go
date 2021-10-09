// This file contain all the functions related to user
// 1.Getting User information
// 2.Creating User
// 3.Deleting User

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vishal206/golang-mongo-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

//In GetUser,this function will get the user from mongoDB with respective the entered user id.

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	userInfo := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&userInfo); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

//In CreateUser, this function will create and store a new user in mongoDb.

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userInfo := models.User{}

	json.NewDecoder(r.Body).Decode(&userInfo)

	userInfo.Id = bson.NewObjectId()

	uc.session.DB("mongo-golang").C("users").Insert(userInfo)

	uj, err := json.Marshal(userInfo)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)

}

//In DeleteUser, this function will delete user in mongoDb with respect to the given user id.

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
