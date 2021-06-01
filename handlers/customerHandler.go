package handlers

import (
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lorezi/golang-bank-app/ports"
	"github.com/lorezi/golang-bank-app/utils"
)

// inject custom service into customer handler
type CustomerHandler struct {
	CustomerService ports.CustomerService
}

func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.CustomerService.GetAllCustomers(status)

	if err != nil {
		utils.Response(w, err.Code, err.ShowError())
		return
	}

	utils.Response(w, http.StatusOK, customers)

}

func (c *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	paramID := mux.Vars(r)
	customer, err := c.CustomerService.GetCustomer(paramID["customer_id"])

	if err != nil {
		utils.Response(w, err.Code, err.ShowError())
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")

		if err := xml.NewEncoder(w).Encode(customer); err != nil {
			panic(err)
		}
		return

	}

	utils.Response(w, http.StatusOK, customer)

}
