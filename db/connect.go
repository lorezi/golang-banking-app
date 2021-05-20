/*
Database connection setup
*/
package db

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

// connects and returns db client
func Connect() *sqlx.DB {

	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbAddr, dbPort, dbName)
	c, err := sqlx.Open("mysql", ds)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	c.SetConnMaxLifetime(time.Minute * 3)
	c.SetMaxOpenConns(10)
	c.SetMaxIdleConns(10)

	return c

}
