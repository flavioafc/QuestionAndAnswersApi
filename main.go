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

var dao = FaqDAO{}

func init() {
	dao.Server = "localhost"
	dao.Database = "db"
	dao.Connect()
}

// @title FAQ API
// @version 1.0
// @description This is a MVP for Questions and Answers for https://www.nuorder.com/ page
// @termsOfService https://www.nuorder.com/
// @contact.name NuOrder API Support
// @contact.email flavio.costa@ecore.com.br
// @license.name Apache 2.0
// @license.url https://www.nuorder.com/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/faq", faqRouter.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/faq/{id}", faqRouter.GetByID).Methods("GET")
	router.HandleFunc("/api/v1/faq", faqRouter.Create).Methods("POST")
	router.HandleFunc("/api/v1/faq/{id}", faqRouter.Update).Methods("PUT")
	router.HandleFunc("/api/v1/faq/{id}", faqRouter.Delete).Methods("DELETE")
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, router))
}
