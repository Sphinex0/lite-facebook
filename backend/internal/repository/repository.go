package repository

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

type Database struct {
	Db *sql.DB
}

var db *sql.DB

func OpenDb() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", "./forum.db")
	return db, err
}