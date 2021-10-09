// This file contain all the functions related to posts
// 1.Getting Posts information
// 2.Creating Posts
// 3.Deleting Posts

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

//In GetPost,this function will get the post from mongoDB with respective the entered post id.

func (uc PostController) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	postInfo := models.Posts{}

	if err := uc.session.DB("mongo-golang").C("posts").FindId(oid).One(&postInfo); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(postInfo)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

//In CreatePost, this function will create and store a new Post in mongoDb.

func (uc PostController) CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	postInfo := models.Posts{}

	json.NewDecoder(r.Body).Decode(&postInfo)

	postInfo.Id = bson.NewObjectId()
	postInfo.PostedTimestamp = time.Now().Format(time.ANSIC)

	uc.session.DB("mongo-golang").C("posts").Insert(postInfo)

	uj, err := json.Marshal(postInfo)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

//In DeletePost, this function will delete post in mongoDb with respect to the given post id.

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
