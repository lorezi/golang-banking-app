package dto

import (
	"strings"

	"github.com/lorezi/golang-bank-app/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"-"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {

	if r.Amount < 5000 {
		return errs.ValidationError("To open a new account you need a minimum of 5000.00", "fail")
	}

	if strings.ToLower(r.AccountType) != "savings" && strings.ToLower(r.AccountType) != "current" {
		return errs.ValidationError("Account type should be savings or current", "fail")
	}

	return nil
}
