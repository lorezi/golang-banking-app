package domain

type Customer struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zip_code"`
	DateofBirth string
	Status      bool
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
