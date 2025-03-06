package repository

import (
	"errors"
	"time"

	"social-network/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (data *Database) CheckMailAndPaswdvalidity(email string, Password string) (int, error) {
	var dbpswd string
	var usrId int
	err := data.Db.QueryRow("SELECT id, password FROM users WHERE email=?", email).Scan(&usrId, &dbpswd)
	if err != nil {
		return 0, errors.New("invalide coredentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbpswd), []byte(Password))
	if err != nil {
		return 0, errors.New("incorrect password")
	}
	return usrId, nil
}

func (database *Database) UpdateUuid(uuid string, userId int) error {
	expire := time.Now().Add(time.Duration(time.Now().Local().Year()))
	_, err := database.Db.Exec("INSERT INTO sessions (user-id = ? uuid = ?, expired_at = ?) VALUES(?,?,?)", userId, uuid, expire)
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
	res, err := database.Db.Exec("INSERT INTO users (Nickname, datebirth, firstName, lastName, email, password, avatar, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		user.Nickname, user.Dob, user.First_Name, user.Last_Name, user.Email, user.Password, user.Image, time.Now())
	if err != nil {
		return err
	}

	usrid, err := res.LastInsertId()
	if err != nil {
		return err
	}

	_, err = database.Db.Exec("INSERT INTO sessions (uuid, user_id, session_exp) VALUES (?,?,?)", user.Uuid, usrid, time.Now().AddDate(1, 0, 0))
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) DeleteCookieFromdb(uuid string) error {
	_, err := database.Db.Exec("DELETE FROM sessions WHERE uuid = ?", uuid)
	if err != nil {
		return err
	}
	return nil
}
