package handlers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/lorezi/golang-bank-app/ports"
)

type CustomerHandlers struct {
	CustomerService ports.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, _ := ch.CustomerService.GetAllCustomers(status)

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

	w.Header().Add("Content-Type", "application/json")
	// encode struct to json
	json.NewEncoder(w).Encode(customers)

}
