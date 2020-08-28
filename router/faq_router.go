package faqrouter

import (
	"encoding/json"
	"net/http"

	. "github.com/flavioafc/go-question-and-answers/dao"
	. "github.com/flavioafc/go-question-and-answers/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var dao = FaqDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GetAll godoc
// @Summary Get a list to all questions and answers from QA API
// @Description Get of all Questions and Answers
// @Tags faq
// @Accept  json
// @Produce  json
// @Success 200 {array} Faq
// @Router /faq [get]
//GetAll list all questions and answers from QA API
func GetAll(w http.ResponseWriter, r *http.Request) {
	faq, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, faq)
}

//GetByID retrieve the question and answer by Id
func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	faq, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid QA ID")
		return
	}
	respondWithJSON(w, http.StatusOK, faq)
}

// Create godoc
// @Summary Create a new Question and Answer
// @Description Create a new Question and Answer with the input paylod
// @Tags faq
// @Accept  json
// @Produce  json
// @Param faq body Faq true "Create"
// @Success 200 {object} Faq
// @Router /faq [post]
//Create method insert in the mongo database a new question and answer
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var faq Faq

	if err := json.NewDecoder(r.Body).Decode(&faq); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	faq.ID = bson.NewObjectId()
	if err := dao.Create(faq); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, faq)
}

//Update have to update the question
func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	var faq Faq
	if err := json.NewDecoder(r.Body).Decode(&faq); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := dao.Update(params["id"], faq); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": faq.ID.String() + ", was successful updated!"})
}

//Delete must delete the question
func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
