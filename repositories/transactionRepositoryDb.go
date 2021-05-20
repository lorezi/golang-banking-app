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

	var err error
	// starting the database transaction block
	tx, err := tr.client.Begin()
	if err != nil {
		logger.Error("error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.UnExpectedError("Unexpected database error", "error")
	}

	// inserting bank account transactions
	insert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values(?,?,?,?)"

	res, err := tr.client.Exec(insert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if err != nil {
		logger.Error("error while processing a new transaction: " + err.Error())
		return nil, errs.UnExpectedError("unexpected error from the database", "error")
	}

	// updating account balance
	// update account for withdrawal action
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	}

	// update account for deposit action
	if !t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	// rollback and changes from both tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("error while saving transaction: " + err.Error())
		return nil, errs.UnExpectedError("Unexpected database error", "error")
	}

	// commit successful transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("error while commiting transaction for bank account: " + err.Error())
		return nil, errs.UnExpectedError("Unexpected database error", "error")
	}

	transactionId, err := res.LastInsertId()
	if err != nil {
		logger.Error("error while getting the last transaction id for new transaction " + err.Error())
		return nil, errs.UnExpectedError("unexpected error from the database", "error")
	}

	// getting the latest account information from the accounts table
	acct, appErr := tr.FindBy(t.AccountId)
	if err != nil {
		return nil, appErr
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// updating the transaction with the latest balance
	t.Amount = acct.Amount

	return &t, nil

}

func (tr TransactionRepositoryDb) FindBy(accountId string) (*domain.Account, *errs.AppError) {
	accountRec := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	account := domain.Account{}
	err := tr.client.Get(&account, accountRec, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.UnExpectedError("Unexpected database error", "error")
	}
	return &account, nil
}
