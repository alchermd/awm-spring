package data

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB
var err error

func init() {
	prepareDB()
}

// Succeeding cals to New will return the same instance of the database
func New() (*sql.DB, error) {
	if db == nil {
		prepareDB()
	}

	return db, err
}

func prepareDB() {
	dns := os.Getenv("DB_DATABASE")
	if dns == "" {
		err = errors.New("$DB_DATABASE must be set.")
		return
	}

	db, err = sql.Open("mysql", dns)
	if err != nil {
		return
	}
}
