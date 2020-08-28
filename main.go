package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/flavioafc/go-question-and-answers/dao"
	faqRouter "github.com/flavioafc/go-question-and-answers/router"
	"github.com/gorilla/mux"
)

var dao = FaqDAO{}

func init() {
	dao.Server = "localhost"
	dao.Database = "db"
	dao.Connect()
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/faq", faqRouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/faq/{id}", faqRouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/faq", faqRouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/faq/{id}", faqRouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/faq/{id}", faqRouter.Delete).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
