package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawal(t *testing.T) {
	// Arrange
	req := TransactionRequest{
		TransactionType: "invalid transaction type",
	}

	// Act
	appErr := req.Validate()

	// Assert for invalid message and code
	if appErr.Message != "The transaction type must be withdrawal or deposit" {
		t.Error("Invalid message while testing type")
	}

	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid http code while testing type")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	req := TransactionRequest{
		TransactionType: DEPOSIT,
		Amount:          -100,
	}

	appErr := req.Validate()

	if appErr.Message != "amount cannot be a negative value" {
		t.Error("Invalid message while testing type")
	}

	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid http code while testing type")
	}
}
