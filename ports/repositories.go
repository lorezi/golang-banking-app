package ports

import "github.com/lorezi/golang-bank-app/domain"

type CustomerRepository interface {
	FindAll() ([]domain.Customer, error)
}
