package dao

import (
	"log"

	. "github.com/flavioafc/go-question-and-answers/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FaqDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "Faq"
)

func (m *FaqDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *FaqDAO) GetAll() ([]Faq, error) {
	var faqs []Faq
	err := db.C(COLLECTION).Find(bson.M{}).All(&faqs)
	return faqs, err
}

func (m *FaqDAO) GetByID(id string) (Faq, error) {
	var faq Faq
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&faq)
	return faq, err
}

func (m *FaqDAO) Create(faq Faq) error {
	err := db.C(COLLECTION).Insert(&faq)
	return err
}

func (m *FaqDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *FaqDAO) Update(id string, faq Faq) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &faq)
	return err
}
