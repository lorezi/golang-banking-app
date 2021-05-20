package service

import (
	"time"

	"github.com/lorezi/golang-bank-app/domain"
	"github.com/lorezi/golang-bank-app/dto"
	"github.com/lorezi/golang-bank-app/errs"
	"github.com/lorezi/golang-bank-app/ports"
)

type DefaultTransactionService struct {
	repo ports.TransactionRepository
}

func NewTransactionService(repo ports.TransactionRepository) *DefaultTransactionService {
	return &DefaultTransactionService{repo}
}

func (s DefaultTransactionService) CreateTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {

	// validate request
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	t := domain.Transaction{
		TransactionId:   "",
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newTrans, err := s.repo.Save(t)
	if err != nil {
		return nil, err
	}

	res := newTrans.DomainToDto()

	return res, nil

}
