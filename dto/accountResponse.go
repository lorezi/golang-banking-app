package dto

type AccountResponse struct {
	AccountId   string  `json:"account_id"`
	CustomerId  string  `json:"customer_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      bool    `json:"status"`
}
