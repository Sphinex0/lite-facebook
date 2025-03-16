package repository

import (
	"database/sql"
	"fmt"
	"time"

	"social-network/internal/models"
	utils "social-network/pkg"

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
	db, err := sql.Open("sqlite3", dbPath+"?_foreign_keys=1")
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

func GetUserByUuid(db *sql.DB, uuid uuid.UUID) (user models.UserInfo, err error) {
	if err = db.QueryRow("SELECT u.id,first_name,last_name,nickname,image FROM users u join sessions s on u.id = s.user_id  WHERE session_id = ? AND session_exp > ?", uuid.String(), time.Now().Unix()).Scan(&user.ID, &user.First_Name, &user.Last_Name, &user.Nickname, &user.Image); err != nil {
		return
	}
	return
}

func (data *Database) GetUserByID(id int) (user models.UserInfo, err error) {
	err = data.Db.QueryRow("SELECT u.id,first_name,last_name,nickname,image FROM users u  WHERE id = ? ", id).Scan(utils.GetScanFields(&user)...)
	return
}
