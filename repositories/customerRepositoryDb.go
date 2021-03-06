package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/lorezi/golang-bank-app/domain"
	"github.com/lorezi/golang-bank-app/errs"
	"github.com/lorezi/golang-bank-app/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (c *CustomerRepositoryDb) FindAll(status string) ([]domain.Customer, *errs.AppError) {

	sc := []domain.Customer{}
	var err error

	allQry := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	if status == "" {
		err = c.client.Select(&sc, allQry)
	}

	if status == "active" {
		allQry = allQry + " where status = 1"
		err = c.client.Select(&sc, allQry)
	}

	if status == "inactive" {
		allQry = allQry + " where status = 0"
		err = c.client.Select(&sc, allQry)
	}

	if err != nil {
		logger.Error("Error while scanning customers" + err.Error())
		return nil, errs.UnExpectedError("unexpected database error", "error")
	}

	return sc, nil
}

func (c *CustomerRepositoryDb) GetById(id string) (*domain.Customer, *errs.AppError) {
	qry := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	customer := &domain.Customer{}

	err := c.client.Get(customer, qry, id)
	if err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("customer with id: %v not found", id)
			return nil, errs.NotFoundError(msg, "fails")
		}

		logger.Error("Error while scanning customers " + err.Error())

		return nil, errs.UnExpectedError("unexpected database error", "error")
	}

	return customer, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) *CustomerRepositoryDb {

	return &CustomerRepositoryDb{client: dbClient}
}
