package dto

type TransactionRequest struct {
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}
