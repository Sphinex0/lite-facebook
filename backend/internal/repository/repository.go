package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

type Database struct {
	Db *sql.DB
}

const dbPath = "internal/repository/forum.db"

func OpenDb() (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlite3", dbPath)
	return db, err
}

func ApplyMigrations(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "pkg/migrations/sqlite",
	}

	_, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("error while executing the migration %v", err)
	}
	fmt.Println("apply Migration Successfully!")
	return nil
}
