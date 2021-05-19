package handlers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lorezi/golang-bank-app/dto"
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

func (c *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	paramID := mux.Vars(r)
	customer, err := c.CustomerService.GetCustomer(paramID["customer_id"])

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		r := dto.Result{
			Status:  "fail",
			Message: "record not found",
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(r)
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	}

	w.Header().Add("Content-Type", "application/json")

	// encode struct to json
	json.NewEncoder(w).Encode(customer)

}
