package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	Zipcode string
}

func greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World!")
}

func getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	sc := []Customer{
		{
			Name: "John Doe", City: "New York", Zipcode: "1100034",
		},
		{
			Name: "Jane Doe", City: "Silicon Valley", Zipcode: "8900001",
		},
	}

	json.NewEncoder(rw).Encode(sc)
}

func main() {
	// defining routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting serve
	log.Fatal(http.ListenAndServe(":8000", nil))
}
