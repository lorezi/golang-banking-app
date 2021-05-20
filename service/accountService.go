package service

import (
	"time"

	"github.com/lorezi/golang-bank-app/domain"
	"github.com/lorezi/golang-bank-app/dto"
	"github.com/lorezi/golang-bank-app/errs"
	"github.com/lorezi/golang-bank-app/ports"
)

type DefaultAccountService struct {
	repo ports.AccountRepository
}

func NewAccountService(repo ports.AccountRepository) *DefaultAccountService {
	return &DefaultAccountService{repo}
}

func (s DefaultAccountService) CreateAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	// validate request
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// transform dto to domain
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      true,
	}

	newAcct, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	res := newAcct.DomainToDto()
	return res, nil
}
