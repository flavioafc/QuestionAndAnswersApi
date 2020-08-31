package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/flavioafc/go-question-and-answers/dao"
	_ "github.com/flavioafc/go-question-and-answers/docs"
	faqRouter "github.com/flavioafc/go-question-and-answers/router"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	Router         *mux.Router
	Dao            DAO
	QuestionRouter *faqRouter.QuestionRouter
	AnswerRouter   *faqRouter.AnswerRouter
}

func (a *App) Initialize(server, database string) {
	a.Dao.Server = server
	a.Dao.Database = database
	a.Dao.Connect()

	a.initializeRoutes()
}

func (a *App) initializeRoutes() {

	a.Router = mux.NewRouter()
	//Questions

	a.Router.HandleFunc("/api/v1/question", a.QuestionRouter.GetAll).Methods("GET")
	a.Router.HandleFunc("/api/v1/question/{id}", a.QuestionRouter.GetByID).Methods("GET")
	a.Router.HandleFunc("/api/v1/question", a.QuestionRouter.Create).Methods("POST")
	a.Router.HandleFunc("/api/v1/question/{id}", a.QuestionRouter.Update).Methods("PUT")
	a.Router.HandleFunc("/api/v1/question/{id}", a.QuestionRouter.Delete).Methods("DELETE")

	//Answers
	a.Router.HandleFunc("/api/v1/answer", a.AnswerRouter.Create).Methods("POST")
	a.Router.HandleFunc("/api/v1/answer/{idquestion}", a.AnswerRouter.GetAll).Methods("GET")
	a.Router.HandleFunc("/api/v1/answer/{id}", a.AnswerRouter.GetByID).Methods("GET")
	a.Router.HandleFunc("/api/v1/answer/{id}", a.AnswerRouter.Update).Methods("PUT")
	a.Router.HandleFunc("/api/v1/answer/{id}", a.AnswerRouter.Delete).Methods("DELETE")

	//Swagger http://localhost:3000/swagger/index.html
	a.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}

func (a *App) Run(port string) {

	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, a.Router))

}
