package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// created multiplexer
	router := mux.NewRouter()

	// defining routes
	router.HandleFunc("/greet", Greet).Methods("GET")
	router.HandleFunc("/customers", GetAllCustomers).Methods("GET")
	// allow customer id with only alpha numeric and underscore character
	router.HandleFunc("/customers/{customer_id:[a-zA-Z0-9_]+}", GetCustomer).Methods("POST")
	router.HandleFunc("/customers", CreateCustomer).Methods("POST")

	// starting serve
	log.Fatal(http.ListenAndServe(":8000", router))
}
