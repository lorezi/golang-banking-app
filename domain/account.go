package domain

import "github.com/lorezi/golang-bank-app/dto"

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      bool    `db:"status"`
}

func (c Account) DomainToDto() *dto.NewAccountResponse {

	return &dto.NewAccountResponse{
		AccountId: c.AccountId,
	}

}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount > amount
}
