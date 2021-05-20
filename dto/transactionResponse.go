package dto

type TransactionResponse struct {
	TransactionId string  `json:"transaction_id"`
	Balance       float64 `json:"balance"`
}
