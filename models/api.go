package models

import "gopkg.in/mgo.v2/bson"

type API struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Type  string
	Count uint64
}
