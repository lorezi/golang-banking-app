package repositories

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/lorezi/golang-bank-app/domain"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (s *CustomerRepositoryDb) FindAll(status string) ([]domain.Customer, error) {

	allQry := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	if status == "active" {
		allQry = allQry + " where status = 1"
	}

	if status == "inactive" {
		allQry = allQry + " where status = 0"
	}

	rows, err := s.client.Query(allQry)

	if err != nil {
		log.Println("Error while querying customer table" + err.Error())
		return nil, err
	}

	sc := []domain.Customer{}
	for rows.Next() {
		c := &domain.Customer{}
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers" + err.Error())
			return nil, err
		}

		sc = append(sc, *c)
	}

	return sc, nil
}

func (s *CustomerRepositoryDb) GetById(id string) (*domain.Customer, error) {
	qry := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := s.client.QueryRow(qry, id)

	c := &domain.Customer{}
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		log.Println("Error while scanning customers" + err.Error())
		return nil, err
	}

	return c, nil
}

func NewCustomerRepositoryDb() *CustomerRepositoryDb {
	c, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	c.SetConnMaxLifetime(time.Minute * 3)
	c.SetMaxOpenConns(10)
	c.SetMaxIdleConns(10)

	return &CustomerRepositoryDb{client: c}
}
