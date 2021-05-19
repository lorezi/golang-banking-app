package domain

type Customer struct {
	Id          string `json:"customer_id" xml:"customer_id" db:"customer_id"`
	Name        string `json:"name" xml:"name" db:"name"`
	City        string `json:"city" xml:"city" db:"city"`
	Zipcode     string `json:"zip_code" xml:"zip_code" db:"zipcode"`
	DateofBirth string `json:"date_of_birth" xml:"date_of_birth" db:"date_of_birth"`
	Status      bool   `json:"status" xml:"status" db:"status"`
}
