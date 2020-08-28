package models

import "gopkg.in/mgo.v2/bson"

type Faq struct {
	ID       bson.ObjectId `bson:"_id" json:"id" example:"5f484f697ee3881a0ca9a037"`
	Question string        `bson:"question" json:"question" example:"How can I remove an item?"`
	Answer   string        `bson:"answer" json:"answer" example:"To remove an item click on RemoveItem button"`
}

type FaqRequest struct {
	Question string `bson:"question" json:"question" example:"How can I remove an item?"`
	Answer   string `bson:"answer" json:"answer" example:"To remove an item click on RemoveItem button"`
}
