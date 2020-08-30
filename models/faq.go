package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Question struct {
	ID    bson.ObjectId `bson:"_id" json:"id" example:"5f484f697ee3881a0ca9a037"`
	Text  string        `bson:"text" json:"text" example:"How can I remove an item?"`
	Image string        `bson:"image" json:"image" example:"http://nuorder.com/images/image.jpg"`
	Likes int           `bson:"likes" json:"likes" example:"30"`
}
type Answer struct {
	ID     bson.ObjectId `bson:"_id" json:"id" example:"7f484f697ee1283a0ca5a028"`
	Parent bson.ObjectId `bson:"parent" json:"parent" example:"5f484f697ee3881a0ca9a037"`
	Text   string        `bson:"text" json:"text" example:"How can I remove an item?"`
	Likes  int           `bson:"likes" json:"likes" example:"10"`
	Answer []*Answer     `bson:"answer" json:"answer,omitempty" example:
	"{
		"_id" : ObjectId("5f4a6a037ee388453cc5aa4b"),
		"parent" : ObjectId("5f4c120c8413c9c01f002e3a"),
		"text" : "My answer for your response  is",
		"likes" : 0
	}"`
}

type FaqRequest struct {
	Question string `bson:"question" json:"question" example:"How can I remove an item?"`
	Answer   string `bson:"answer" json:"answer" example:"To remove an item click on RemoveItem button"`
}
