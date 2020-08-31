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

var retorno []*Answer

func (m *AnswerDAO) ClearMemory() {
	retorno = nil
}

func (m *AnswerDAO) GetAll(id string) ([]*Answer, error) {

	var faqs []*Answer
	var isChild bool = false
	err := db.C(COLLECTION_ANSWER).Find(bson.M{"parent": bson.ObjectIdHex(id)}).All(&faqs)

	for _, item := range faqs {
		node := &Answer{
			ID:     item.ID,
			Text:   item.Text,
			Root:   item.Root,
			Parent: item.Parent,
			Likes:  item.Likes,
			Answer: item.Answer,
		}

		for i := range retorno {
			if node.Parent == retorno[i].ID {
				retorno[i].Answer = append(retorno[i].Answer, node)
				isChild = true
				m.GetAll(item.ID.Hex())
			} else {
				if len(retorno[i].Answer) > 0 {
					for _, w := range retorno[i].Answer {
						if node.Parent == w.ID {
							w.Answer = append(w.Answer, node)
							isChild = true
						}
					}
				}
			}
		}

		if isChild == false {
			retorno = append(retorno, node)
			m.GetAll(item.ID.Hex())
		}
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
