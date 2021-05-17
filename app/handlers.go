package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/lorezi/golang-bank-app/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

	w.Header().Add("Content-Type", "application/json")
	// encode struct to json
	json.NewEncoder(w).Encode(customers)

}
