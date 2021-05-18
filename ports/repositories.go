package ports

import "github.com/lorezi/golang-bank-app/domain"

type CustomerRepository interface {
	FindAll(status string) ([]domain.Customer, error)
	GetById(id string) (*domain.Customer, error)
}
