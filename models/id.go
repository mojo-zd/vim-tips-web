package models

import "gopkg.in/mgo.v2/bson"

type Identity struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Email    string
	Password []byte
}
