package service

import (
	"github.com/lorezi/golang-bank-app/dto"
	"github.com/lorezi/golang-bank-app/errs"
	"github.com/lorezi/golang-bank-app/ports"
)

type DefaultCustomerService struct {
	repo ports.CustomerRepository
}

func NewCustomerService(repository ports.CustomerRepository) *DefaultCustomerService {
	return &DefaultCustomerService{repo: repository}
}

func (s *DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {

	sc := []dto.CustomerResponse{}
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	for _, v := range customers {
		res := v.DomainToDto()
		sc = append(sc, *res)
	}

	return sc, nil

}

func (s *DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	res := c.DomainToDto()

	return res, nil

}
