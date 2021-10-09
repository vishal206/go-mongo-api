package main

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestPosts(t *testing.T) {
	// t.Errorf("hello")

	id := "61617be95f92d44c288ad519"

	if !bson.IsObjectIdHex(id) {
		t.Errorf("status not found error")
	}
}

func TestUsers(t *testing.T) {
	id := "61615d195f92d458c0153cd1"

	if !bson.IsObjectIdHex(id) {
		t.Errorf("status not found error")
	}

}
