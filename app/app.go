package app

import (
	"log"
	"net/http"
)

func Start() {
	// defining routes
	http.HandleFunc("/greet", Greet)
	http.HandleFunc("/customers", GetAllCustomers)

	// starting serve
	log.Fatal(http.ListenAndServe(":8000", nil))
}
