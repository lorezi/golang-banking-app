package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	sc := []Customer{
		{
			Name: "John Doe", City: "New York", Zipcode: "1100034",
		},
		{
			Name: "Jane Doe", City: "Silicon Valley", Zipcode: "8900001",
		},
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(sc)
}

func main() {
	// defining routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting serve
	log.Fatal(http.ListenAndServe(":8000", nil))
}
