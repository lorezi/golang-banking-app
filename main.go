package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
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

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(sc)
	}

	w.Header().Add("Content-Type", "application/json")
	// encode struct to json
	json.NewEncoder(w).Encode(sc)

}

func main() {
	// defining routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting serve
	log.Fatal(http.ListenAndServe(":8000", nil))
}
