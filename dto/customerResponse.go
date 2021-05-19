package dto

type CustomerResponse struct {
	Id          string `json:"customer_id" xml:"customer_id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zip_code"`
	DateofBirth string `json:"date_of_birth" xml:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}
