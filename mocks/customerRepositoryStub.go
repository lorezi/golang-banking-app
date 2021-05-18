package mocks

import "github.com/lorezi/golang-bank-app/domain"

type CustomerRepositoryStub struct {
	customers []domain.Customer
}

func (s CustomerRepositoryStub) FindAll() ([]domain.Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	sc := []domain.Customer{
		{
			Name: "John Doe", City: "New York", Zipcode: "1100034", DateofBirth: "2000-01-04", Status: true, Id: "100001",
		},
		{
			Name: "Jane Doe", City: "Silicon Valley", Zipcode: "8900001", DateofBirth: "2000-01-08", Status: false, Id: "100034",
		},
	}

	return CustomerRepositoryStub{customers: sc}
}
