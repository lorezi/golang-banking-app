package dto

import (
	"strings"

	"github.com/lorezi/golang-bank-app/errs"
)

const (
	WITHDRAWAL = "withdrawal"
	DEPOSIT    = "deposit"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	CustomerId      string  `string:"customer_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

func (t TransactionRequest) Validate() *errs.AppError {

	if strings.ToLower(t.TransactionType) != WITHDRAWAL && strings.ToLower(t.TransactionType) != DEPOSIT {
		return errs.ValidationError("The transaction type must be withdrawal or deposit", "fail")
	}

	if t.Amount < 0 {
		return errs.ValidationError("amount cannot be a negative value", "fail")
	}

	return nil
}

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {

	return true
}
