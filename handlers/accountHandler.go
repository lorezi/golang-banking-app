package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lorezi/golang-bank-app/dto"
	"github.com/lorezi/golang-bank-app/ports"
	"github.com/lorezi/golang-bank-app/utils"
)

// inject account service into account handler
type AccountHandler struct {
	AccountService ports.AccountService
}

func (ah AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	customerId := mux.Vars(r)["customer_id"]

	req := dto.NewAccountRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Response(w, http.StatusBadRequest, "invalid data 🥵🥵")
		return
	}

	req.CustomerId = customerId

	res, err := ah.AccountService.CreateAccount(req)
	if err != nil {
		utils.Response(w, err.Code, err.ShowError())
		return
	}

	utils.Response(w, http.StatusOK, res)

}
