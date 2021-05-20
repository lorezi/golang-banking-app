package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lorezi/golang-bank-app/dto"
	"github.com/lorezi/golang-bank-app/ports"
	"github.com/lorezi/golang-bank-app/utils"
)

type TransactionHandler struct {
	Service ports.TransactionService
}

func (th TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// retrieve the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	req := dto.TransactionRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Response(w, http.StatusBadRequest, "invalid data 🥵🥵")
		return
	}

	req.AccountId = accountId
	req.CustomerId = customerId

	res, err := th.Service.CreateTransaction(req)
	if err != nil {
		utils.Response(w, err.Code, err.ShowError())
		return
	}

	utils.Response(w, http.StatusOK, res)

}
