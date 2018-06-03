package models

import "gopkg.in/mgo.v2/bson"

type Tips struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	Content string
	Comment string
}

type TipsView struct {
	Id      string
	Content string
	Comment string
}
