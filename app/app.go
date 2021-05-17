package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lorezi/golang-bank-app/domain"
	"github.com/lorezi/golang-bank-app/service"
)

func Start() {

	// wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// created multiplexer
	router := mux.NewRouter()

	// defining routes

	router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	// allow customer id with only alpha numeric and underscore character
	// router.HandleFunc("/customers/{customer_id:[a-zA-Z0-9_]+}", GetCustomer).Methods("POST")
	// router.HandleFunc("/customers", CreateCustomer).Methods("POST")

	// starting serve
	log.Fatal(http.ListenAndServe(":8000", router))
}
