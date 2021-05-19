package ports

import (
	"github.com/lorezi/golang-bank-app/domain"
	"github.com/lorezi/golang-bank-app/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, error)
	GetCustomer(id string) (*domain.Customer, *errs.AppError)
}
