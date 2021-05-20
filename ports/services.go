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
	CreateAccount(*dto.AccountResponse, *errs.AppError)
}
