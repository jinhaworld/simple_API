package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	maxIdleConns = 10
	maxOpenConns = 10
	maxConnLife  = 5 * time.Minute
)

var (
	db *sql.DB
)

type Book struct {
	ID        string `db:"id"`
	Title     string `db:"name"`
	CreatedAt string `db:string`
}

func GetDB(connString string) (*sql.DB, error) {
	if db == nil {
		log.Println("Creating a new connection...")

		d, err := sql.Open("mysql", connString)
		if err != nil {
			return nil, err
		}
		err = d.Ping()
		if err != nil {
			return nil, err
		}

		db = d

		db.SetMaxIdleConns(maxIdleConns)
		db.SetMaxOpenConns(maxOpenConns)
		db.SetConnMaxLifetime(maxConnLife)
	}

	return db, nil
}
