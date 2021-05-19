package ports

import (
	"github.com/lorezi/golang-bank-app/domain"
	"github.com/lorezi/golang-bank-app/errs"
)

type CustomerRepository interface {
	FindAll(status string) ([]domain.Customer, error)
	GetById(id string) (*domain.Customer, *errs.AppError)
}
