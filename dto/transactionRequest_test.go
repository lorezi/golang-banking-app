package dto

import "testing"

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawal(t *testing.T) {
	// Arrange
	req := TransactionRequest{
		TransactionType: "invalid transaction type",
	}

	// Act
	appErr := req.Validate()

	// Assert
	if appErr.Message != "The transaction type must be withdrawal or deposit" {
		t.Error("Invalid message while testing type")
	}
}
