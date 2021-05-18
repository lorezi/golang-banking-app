package repositories

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/lorezi/golang-bank-app/domain"
)

type CustomerRepositoryDb struct {
}

func (s CustomerRepositoryDb) FindAll() ([]domain.Customer, error) {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return nil, nil
}
