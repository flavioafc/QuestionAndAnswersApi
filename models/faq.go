package models

import "gopkg.in/mgo.v2/bson"

type Faq struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Question string        `bson:"question" json:"question"`
	Answer   string        `bson:"answer" json:"answer"`
}
