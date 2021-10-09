//this is the model used for posts information
package models

import "gopkg.in/mgo.v2/bson"

type Posts struct {
	Id              bson.ObjectId `json:"id" bson:"_id"`
	Caption         string        `json:"caption" bson:"caption"`
	Imageurl        string        `json:"imageurl" bson:"imageurl"`
	PostedTimestamp string        `json:"postedtimestamp" bson:"postedtimestamp"`
}
