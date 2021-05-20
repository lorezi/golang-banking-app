package dto

import (
	"strings"

	"github.com/lorezi/golang-bank-app/errs"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

func (t TransactionRequest) Validate() *errs.AppError {

	if strings.ToLower(t.TransactionType) != "withdrawal" && strings.ToLower(t.TransactionType) != "deposit" {
		return errs.ValidationError("The transaction type must be withdrawal or deposit", "fail")
	}

	if t.Amount < 0 {
		return errs.ValidationError("amount cannot be a negative value", "fail")
	}

	return nil
}
