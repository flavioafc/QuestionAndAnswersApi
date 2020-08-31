package router

import (
	"encoding/json"
	"net/http"

	. "github.com/flavioafc/go-question-and-answers/dao"
	. "github.com/flavioafc/go-question-and-answers/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var question_dao = QuestionDAO{}

type QuestionRouter struct {
}

// GetAll godoc
// @Summary Get a list to all questions from the API
// @Description Get a list of all questions
// @Tags Question
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Question
// @Router /api/v1/question [get]
// GetAll list all questions and answers from QA API
func (q *QuestionRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	faq, err := question_dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, faq)
}

// GetByID godoc
// @Summary Get one question and answer item from the API
// @Description Get a question and answer
// @Tags Question
// @Accept  json
// @Produce  json
// @Param id path string true "ObjectId"
// @Success 200 {object} models.Question
// @Router /api/v1/question/{id} [get]
// GetByID retrieve the question and answer by Id
func (q *QuestionRouter) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	faq, err := question_dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Question ID")
		return
	}
	respondWithJSON(w, http.StatusOK, faq)
}

// Create godoc
// @Summary Create a new Question and Answer item
// @Description Create a new Question and Answer with the input paylod
// @Tags Question
// @Accept  json
// @Produce  json
// @Param faq body models.Question true "Create"
// @Success 200 {object} models.Question
// @Router /api/v1/question [post]
// Create method insert in the mongo database a new question and answer
func (q *QuestionRouter) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var faq Question

	if err := json.NewDecoder(r.Body).Decode(&faq); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	faq.ID = bson.NewObjectId()
	if err := question_dao.Create(faq); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, faq)
}

// Update godoc
// @Summary Update a new Question and Answer item
// @Description Update a new Question and Answer with the input paylod
// @Tags Question
// @Accept  json
// @Produce  json
// @Param id path string true "ObjectId"
// @Param faq body models.Question true "Update"
// @Success 200 "ObjectIdHex(id), was successful updated!"
// @Router /api/v1/question/{id} [put]
// Update have to update the question
func (q *QuestionRouter) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	var faq Question
	if err := json.NewDecoder(r.Body).Decode(&faq); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := question_dao.Update(params["id"], faq); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": faq.ID.String() + ", was successful updated!"})
}

// Delete godoc
// @Summary Delete one question and  all answers items from the API
// @Description Delete a question and all the answers
// @Tags Question
// @Accept  json
// @Produce  json
// @Param id path string true "ObjectId"
// @Success 200 "result: success"
// @Router /api/v1/question/{id} [delete]
// Delete must delete the question
func (q *QuestionRouter) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := question_dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err := answer_dao.DeleteRelatedByRoot(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
