package repositories

import (
	"database/sql"
	"fmt"
	"os"

	"time"

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

func NewCustomerRepositoryDb() *CustomerRepositoryDb {
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	c, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbAddr, dbPort, dbName))
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	c.SetConnMaxLifetime(time.Minute * 3)
	c.SetMaxOpenConns(10)
	c.SetMaxIdleConns(10)

	return &CustomerRepositoryDb{client: c}
}
