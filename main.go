package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/vishal206/golang-mongo-api/controllers"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/users/:id", uc.GetUser)
	r.POST("/users", uc.CreateUser)
	r.DELETE("/users/:id", uc.DeleteUser)
	pc := controllers.NewPostController(getSession())
	r.GET("/posts/:id", pc.GetPost)
	r.POST("/posts", pc.CreatePost)
	r.DELETE("/posts/:id", pc.DeletePost)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
