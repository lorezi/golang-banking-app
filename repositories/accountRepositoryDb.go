package repositories

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/lorezi/golang-bank-app/domain"
	"github.com/lorezi/golang-bank-app/errs"
	"github.com/lorezi/golang-bank-app/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (ar AccountRepositoryDb) Save(a domain.Account) (*domain.Account, *errs.AppError) {
	insert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values(?,?,?,?,?)"
	res, err := ar.client.Exec(insert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("error while creating a new account: " + err.Error())
		return nil, errs.UnExpectedError("unexpected error from the database", "error")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("error while getting last insert id for new account " + err.Error())
		return nil, errs.UnExpectedError("unexpected error from the database", "error")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
