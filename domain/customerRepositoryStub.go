package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	sc := []Customer{
		{
			Name: "John Doe", City: "New York", Zipcode: "1100034", DateofBirth: "2000-01-04", Status: true, Id: "100001",
		},
		{
			Name: "Jane Doe", City: "Silicon Valley", Zipcode: "8900001", DateofBirth: "2000-01-08", Status: false, Id: "100034",
		},
	}

	return CustomerRepositoryStub{customers: sc}
}
