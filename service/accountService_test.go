package service

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/lorezi/golang-bank-app/domain"
	"github.com/lorezi/golang-bank-app/dto"
	"github.com/lorezi/golang-bank-app/errs"
	"github.com/lorezi/golang-bank-app/mocks"
)

var mk *mocks.MockAccountRepository
var s *DefaultAccountService

// mock and service global setup
func setup(t *testing.T) func() {

	// mock setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mk = mocks.NewMockAccountRepository(ctrl)

	// service setup
	s = NewAccountService(mk)

	return func() {
		defer ctrl.Finish()
	}
}

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

func Test_should_return_server_side_error_when_the_new_account_is_not_created(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "savings",
		Amount:      6000,
	}
	acct := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      true,
	}

	mk.EXPECT().Save(acct).Return(nil, errs.UnExpectedError("database error", "fail"))

	// Act
	_, err := s.CreateAccount(req)

	// Assert
	if err == nil {
		t.Error("failed while testing the new account creation")
	}
}

func Test_should_return_new_account_response_when_account_is_created_successfully(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "savings",
		Amount:      6000,
	}
	acct := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      true,
	}

	acctWithId := acct
	acctWithId.AccountId = "200"

	mk.EXPECT().Save(acct).Return(&acctWithId, nil)

	// Act
	newAcct, err := s.CreateAccount(req)

	// Assert
	if err != nil {
		t.Error("Test failed while creating new account")
	}

	if newAcct.AccountId != acctWithId.AccountId {
		t.Error("Test failed while matching new account id")
	}
}
