package service

import (
	"testing"

	"github.com/lorezi/golang-bank-app/dto"
)

func Test_should_return_validation_error_response_when_the_new_account_request_fails_validation(t *testing.T) {
	// Arrange
	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0,
	}

	s := NewAccountService(nil)

	// Act
	_, err := s.CreateAccount(req)

	// Assert
	if err == nil {
		t.Error("failed while testing the new account validation")
	}
}
