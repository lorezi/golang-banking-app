package service

import (
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
	if err := req.Validate(); err != nil {
		return nil, err
	}

	a := domain.DtoToDomain(req)

	newAcct, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	return newAcct.DomainToDto(), nil
}
