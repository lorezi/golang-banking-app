package dto

import (
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

	if !t.IsTransactionTypeWithdrawal() && !t.IsTransactionTypeDeposit() {
		return errs.ValidationError("The transaction type must be withdrawal or deposit", "fail")
	}

	if t.Amount < 0 {
		return errs.ValidationError("amount cannot be a negative value", "fail")
	}

	return nil
}

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t TransactionRequest) IsTransactionTypeDeposit() bool {
	return t.TransactionType == DEPOSIT
}
