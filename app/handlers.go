package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/lorezi/golang-bank-app/models"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	sc := []models.Customer{
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
