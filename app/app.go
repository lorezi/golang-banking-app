package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	addr := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), router))
}
