package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lorezi/golang-bank-app/handlers"
	"github.com/lorezi/golang-bank-app/repositories"
	"github.com/lorezi/golang-bank-app/service"
)

func Start() {

	// Testing
	// ch := handlers.CustomerHandlers{
	// 	CustomerService: service.NewCustomerService(mocks.NewCustomerRepositoryStub()),
	// }

	// wiring
	ch := handlers.CustomerHandlers{
		CustomerService: service.NewCustomerService(repositories.NewCustomerRepositoryDb()),
	}

	// created multiplexer
	router := mux.NewRouter()

	// defining routes

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods("GET")
	// allow customer id with only alpha numeric and underscore character
	router.HandleFunc("/customers/{customer_id:[a-zA-Z0-9_]+}", ch.GetCustomer).Methods("GET")
	// router.HandleFunc("/customers", ch.GetAllCustomers).Methods("GET")

	// starting serve
	log.Fatal(http.ListenAndServe(":8000", router))
}
