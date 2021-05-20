package domain

import "github.com/lorezi/golang-bank-app/dto"

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) DomainToDto() *dto.TransactionResponse {

	return &dto.TransactionResponse{
		TransactionId: t.TransactionId,
		Balance:       t.Amount,
	}

}
