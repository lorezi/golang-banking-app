package domain

import "github.com/lorezi/golang-bank-app/dto"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateofBirth string `db:"date_of_birth"`
	Status      bool   `db:"status"`
}

func (c Customer) statusAsText() string {
	status := "inactive"

	if c.Status {
		status = "active"
	}

	return status
}

func (c Customer) DomainToDto() *dto.CustomerResponse {

	return &dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		DateofBirth: c.DateofBirth,
		Zipcode:     c.Zipcode,
		Status:      c.statusAsText(),
		City:        c.City,
	}

}
