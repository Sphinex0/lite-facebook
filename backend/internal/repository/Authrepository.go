package repository

import (
	"errors"
	"social-network/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (data *Database) CheckMailAndPaswdvalidity(email string, Password string) error {
	var dbpswd string
	err := data.Db.QueryRow("SELECT password FROM users WHERE email=?", email).Scan(&dbpswd)
	if err != nil {
		return errors.New("invalide coredentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbpswd), []byte(Password))
	if err != nil {
		return errors.New("incorrect password")
	}
	return nil
}

func (database *Database) UpdateUuid(uuid, email string) error {
	expire := time.Now().Add(time.Duration(time.Now().Local().Year()))
	_, err := database.Db.Exec("UPDATE users SET uuid = ?, expired_at = ? WHERE email = ?", uuid, expire, email)
	return err
}

func (Database *Database) GetUser(uid string) (int, error) {
	var id int
	err := Database.Db.QueryRow("SELECT id FROM users WHERE uuid = ?", uid).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (database *Database) CheckIfUserExists(email string) bool {
	var uname string
	err := database.Db.QueryRow("SELECT Nickname FROM users WHERE email = ?", email).Scan(&uname)
	return err == nil
}

func (database *Database) InsertUser(user models.User) error {
	_, err := database.Db.Exec("INSERT INTO user (Nickname, Age, First_Name, Last_Name, email, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		user.Nickname, user.Age, user.First_Name, user.Last_Name, user.Email, user.Password)
	return err
}