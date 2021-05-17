package app

import (
	"log"
	"net/http"
)

func Start() {

	// created multiplexer
	mux := http.NewServeMux()

	// defining routes
	mux.HandleFunc("/greet", Greet)
	mux.HandleFunc("/customers", GetAllCustomers)

	// starting serve
	log.Fatal(http.ListenAndServe(":8000", mux))
}
