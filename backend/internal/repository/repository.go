package repository

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Db *sql.DB
}

var db *sql.DB

const dbPath = "internal/repository/forum.db"

func OpenDb() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	return db, err
}

func ApplyMigrations() error {
	// get the currente directory path name
	wd, err := os.Getwd()
	fmt.Println(wd)
	if err != nil {
		return fmt.Errorf("can't get the rooted path name %w", err)
	}

	// Making the migration path wre the migrations are in
	migrationsPath := "file://" + filepath.Join(wd, "..","backend","pkg", "migrations","sqlite")

	// return and instance on migration to work with
	m, err := migrate.New(migrationsPath, "sqlite3://"+dbPath)
	if err != nil {
		return fmt.Errorf("error while loading migrations: %w", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("no migration changes")
			return nil
		}
		return fmt.Errorf("error while migrating the application: %w", err)
	}

	fmt.Println("migration succesfully !")
	return nil
}