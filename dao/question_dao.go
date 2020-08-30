package dao

import (
	. "github.com/flavioafc/go-question-and-answers/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	COLLECTION_QUESTION = "question"
)

type QuestionDAO struct{}

func (m *QuestionDAO) GetAll() ([]Question, error) {
	var faqs []Question
	err := db.C(COLLECTION_QUESTION).Find(bson.M{}).All(&faqs)
	return faqs, err
}

func (m *QuestionDAO) GetByID(id string) (Question, error) {
	var faq Question
	err := db.C(COLLECTION_QUESTION).FindId(bson.ObjectIdHex(id)).One(&faq)
	return faq, err
}

func (m *QuestionDAO) Create(faq Question) error {
	err := db.C(COLLECTION_QUESTION).Insert(&faq)
	return err
}

func (m *QuestionDAO) Delete(id string) error {
	err := db.C(COLLECTION_QUESTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *QuestionDAO) Update(id string, faq Question) error {
	err := db.C(COLLECTION_QUESTION).UpdateId(bson.ObjectIdHex(id), &faq)
	return err
}
