package dao

import (
	. "github.com/flavioafc/go-question-and-answers/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AnswerDAO struct {
}

const (
	COLLECTION_ANSWER = "answer"
)

func (m *AnswerDAO) GetAll(id string) ([]*Answer, error) {
	var retorno []*Answer
	var faqs []*Answer

	err := db.C(COLLECTION_ANSWER).Find(bson.M{"parent": bson.ObjectIdHex(id)}).All(&faqs)
	for _, item := range faqs {
		node := &Answer{}
		node.ID = item.ID
		node.Text = item.Text
		node.Root = item.Root
		node.Parent = item.Parent
		node.Likes = item.Likes
		node.Answer, _ = m.GetAll(item.ID.Hex())
		retorno = append(retorno, node)
	}

	return retorno, err
}

func (m *AnswerDAO) GetByID(id string) (Answer, error) {
	var faq Answer
	err := db.C(COLLECTION_ANSWER).FindId(bson.ObjectIdHex(id)).One(&faq)
	return faq, err
}

func (m *AnswerDAO) Create(faq AnswerRequest) error {
	err := db.C(COLLECTION_ANSWER).Insert(&faq)
	return err
}

func (m *AnswerDAO) Delete(id string) error {
	err := db.C(COLLECTION_ANSWER).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *AnswerDAO) DeleteRelatedByRoot(idRoot string) (*mgo.ChangeInfo, error) {
	result, err := db.C(COLLECTION_ANSWER).RemoveAll(bson.M{"parent": bson.ObjectIdHex(idRoot)})
	return result, err
}

func (m *AnswerDAO) Update(id string, faq Answer) error {
	err := db.C(COLLECTION_ANSWER).UpdateId(bson.ObjectIdHex(id), &faq)
	return err
}
