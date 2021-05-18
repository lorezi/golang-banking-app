package ports

import "github.com/lorezi/golang-bank-app/domain"

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, error)
	GetCustomer(id string) (*domain.Customer, error)
}
