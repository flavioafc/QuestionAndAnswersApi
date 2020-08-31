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

var dao = DAO{}
var questionRouter = faqRouter.QuestionRouter{}
var answerRouter = faqRouter.AnswerRouter{}

func init() {
	dao.Server = "localhost"
	dao.Database = "db"
	dao.Connect()
}

// @title FAQ RESTful Service API
// @version 1.0
// @description This is a MVP for Questions and Answers for https://www.nuorder.com/ page.
// @description Install MongoDB to test this service or use Docker executing Compose up in docker-compose.yml in the project folder check the CLI for reference in https://docs.docker.com/compose/reference/up/
// @description If executing this service by VS Code, just click on right button on docker-compose.yml and Compose Up
// @termsOfService https://www.nuorder.com/
// @contact.name NuOrder API Support
// @contact.email flavio.costa@ecore.com.br
// @license.name Apache 2.0
// @license.url https://www.nuorder.com/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {

	router := mux.NewRouter()

	//Questions
	router.HandleFunc("/api/v1/question", questionRouter.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/question/{id}", questionRouter.GetByID).Methods("GET")
	router.HandleFunc("/api/v1/question", questionRouter.Create).Methods("POST")
	router.HandleFunc("/api/v1/question/{id}", questionRouter.Update).Methods("PUT")
	router.HandleFunc("/api/v1/question/{id}", questionRouter.Delete).Methods("DELETE")

	//Answers
	router.HandleFunc("/api/v1/answer", answerRouter.Create).Methods("POST")
	router.HandleFunc("/api/v1/answer/{idquestion}", answerRouter.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/answer/{id}", answerRouter.GetByID).Methods("GET")
	router.HandleFunc("/api/v1/answer/{id}", answerRouter.Update).Methods("PUT")
	router.HandleFunc("/api/v1/answer/{id}", answerRouter.Delete).Methods("DELETE")

	//Swagger
	//Address: http://localhost:3000/swagger/index.html
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, router))
}
