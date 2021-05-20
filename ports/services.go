package ports

import (
	"github.com/lorezi/golang-bank-app/dto"
	"github.com/lorezi/golang-bank-app/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type AccountService interface {
	CreateAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type TransactionService interface {
	CreateTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}
