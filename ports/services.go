package ports

import "github.com/lorezi/golang-bank-app/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}
