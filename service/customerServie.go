package service

import (
	"github.com/lorezi/golang-bank-app/domain"
	"github.com/lorezi/golang-bank-app/ports"
)

type DefaultCustomerService struct {
	repo ports.CustomerRepository
}

func NewCustomerService(repository ports.CustomerRepository) *DefaultCustomerService {
	return &DefaultCustomerService{repo: repository}
}

func (s *DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, error) {

	return s.repo.FindAll(status)
}
