package repositories

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/lorezi/golang-bank-app/domain"
	"github.com/lorezi/golang-bank-app/errs"
	"github.com/lorezi/golang-bank-app/logger"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{client: dbClient}
}

func (tr TransactionRepositoryDb) Save(t domain.Transaction) (*domain.Transaction, *errs.AppError) {

	insert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values(?,?,?,?)"

	res, err := tr.client.Exec(insert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if err != nil {
		logger.Error("error while processing a new transaction: " + err.Error())
		return nil, errs.UnExpectedError("unexpected error from the database", "error")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("error while getting the last insert id for new transaction " + err.Error())
		return nil, errs.UnExpectedError("unexpected error from the database", "error")
	}

	t.TransactionId = strconv.FormatInt(id, 10)

	return &t, nil

}
