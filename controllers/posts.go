package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/vishal206/golang-mongo-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PostController struct {
	session *mgo.Session
}

func NewPostController(s *mgo.Session) *PostController {
	return &PostController{s}
}

func (uc PostController) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.Posts{}

	if err := uc.session.DB("mongo-golang").C("posts").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc PostController) CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.Posts{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()
	u.PostedTimestamp = time.Now().Format(time.ANSIC)

	uc.session.DB("mongo-golang").C("posts").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)

	// fmt.Println("hello world") 616161905f92d43c14800bdd
}

func (uc PostController) DeletePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("posts").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted post", oid, "\n")
}
