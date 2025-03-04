package repository

import (
	"database/sql"
	"fmt"
	"time"

	"social-network/internal/models"

	"github.com/gofrs/uuid/v5"
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

func (data *Database) StoreUser(user models.User) {
	data.Db.Exec("")
}

func (data *Database) StoreSession(user models.User) {
	data.Db.Exec("")
}

func GetUserByUuid(db *sql.DB, uuid uuid.UUID) (user models.User, err error) {
	if err = db.QueryRow("SELECT id,first_name,last_name,nickname,image FROM users WHERE uuid = ? AND uuid_exp > ?", uuid.String(), time.Now().Unix()).Scan(&user.ID, &user.First_Name, &user.Last_Name, &user.Nickname, &user.Image); err != nil {
		return
	}
	return
}
