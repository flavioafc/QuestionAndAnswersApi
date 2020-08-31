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

type AnswerRequest struct {
	ID     bson.ObjectId `bson:"_id" json:"id" example:"7f484f697ee1283a0ca5a028"`
	Root   bson.ObjectId `bson:"root" json:"root" example:"5f484f697ee3881a0ca9a037"`
	Parent bson.ObjectId `bson:"parent" json:"parent" example:"5f484f697ee3881a0ca9a037"`
	Text   string        `bson:"text" json:"text" example:"How can I remove an item?"`
	Likes  int           `bson:"likes" json:"likes" example:"10"`
}

type Answer struct {
	ID     bson.ObjectId `bson:"_id" json:"id" example:"7f484f697ee1283a0ca5a028"`
	Root   bson.ObjectId `bson:"root" json:"root" example:"5f484f697ee3881a0ca9a037"`
	Parent bson.ObjectId `bson:"parent" json:"parent" example:"5f484f697ee3881a0ca9a037"`
	Text   string        `bson:"text" json:"text" example:"How can I remove an item?"`
	Likes  int           `bson:"likes" json:"likes" example:"10"`
	Answer []*Answer     `bson:"answer" json:"answer,omitempty"`
}

type FaqRequest struct {
	Question string `bson:"question" json:"question" example:"How can I remove an item?"`
	Answer   string `bson:"answer" json:"answer" example:"To remove an item click on RemoveItem button"`
}
