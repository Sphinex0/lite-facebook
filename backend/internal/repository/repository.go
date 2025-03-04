package repository

import (
	"database/sql"
	"fmt"
	"social-network/internal/models"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

type Database struct {
	Db *sql.DB
}

const dbPath = "internal/repository/forum.db"

func OpenDb() (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlite3", dbPath+"?_foreing_keys=1")
	if err != nil {
		return db, err
	}
	db.SetMaxOpenConns(10)
	return db, nil
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

func (data *Database) StoreUser(user models.User)  {
	data.Db.Exec("")
}

func (data *Database) StoreSession(user models.User)  {
	data.Db.Exec("")
}