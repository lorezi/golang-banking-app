package domain

import (
	"time"

	"github.com/lorezi/golang-bank-app/dto"
)

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

func DtoToDomain(req dto.NewAccountRequest) Account {
	return Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      true,
	}
}
