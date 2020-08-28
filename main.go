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
// @description This is a MVP for Questions and Answers https://www.nuorder.com/ page
// @termsOfService http://swagger.io/terms/
// @contact.name NuOrder API Support
// @contact.email flavio.costa@swagger.io
// @license.name Apache 2.0
// @license.url https://www.nuorder.com/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/faq", faqRouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/faq/{id}", faqRouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/faq", faqRouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/faq/{id}", faqRouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/faq/{id}", faqRouter.Delete).Methods("DELETE")
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
