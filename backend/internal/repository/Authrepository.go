package repository

import (
	"database/sql"
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

func (database *Database) AddUuid(Uuid string, userId int) error {
	_, err := database.Db.Exec("INSERT INTO sessions (uuid, user_id, session_exp) VALUES (?,?,?)", Uuid, userId, time.Now().AddDate(1, 0, 0))
	if err != nil {
		return err
	}
	return nil
}

func (Database *Database) GetUser(uid string) (int, error) {
	var id int
	err := Database.Db.QueryRow("SELECT id FROM sessions WHERE uuid = ?", uid).Scan(&id)
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

func (database *Database) InsertUser(user models.User, Uuid string) error {
	res, err := database.Db.Exec("INSERT INTO users (Nickname, date_birth, first_name, last_name, email, password, image, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		user.Nickname, user.DateBirth, user.First_Name, user.Last_Name, user.Email, user.Password, user.Image, time.Now())
	if err != nil {
		return err
	}

	usrid, err := res.LastInsertId()
	if err != nil {
		return err
	}

	err = database.AddUuid(Uuid, int(usrid))
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

func CheckIfUserExistsById[T int | string](usrID T, Db *sql.DB) bool {
	var exists bool

	err := Db.QueryRow("SELECT EXISTS(SELECT id FROM users WHERE usrid = ?)", usrID).Scan(&exists)

	return err == nil
}

func CheckGroupIfExistsById(GroupId int, Db *sql.DB) bool {
	var exists bool

	err := Db.QueryRow("SELECT EXISTS(SELECT id FROM groups WHERE id = ?)", GroupId).Scan(&exists)

	return err == nil
}

func (database *Database) CheckExpiredCookie(uid string, date time.Time) bool {
	var expired time.Time
	database.Db.QueryRow("SELECT session_exp FROM sessions WHERE uuid = ?", uid).Scan(&expired)

	return date.Compare(expired) <= -1
}

func (database *Database) GetuserInfo(userId int) (models.UserInfo ,error) {
	var userInfo models.UserInfo
	err := database.Db.QueryRow("SELECT id, Nickname, First_Name, Last_Name, Image FROM users WHERE id = ?", userId).Scan(&userInfo.ID, &userInfo.Nickname, &userInfo.First_Name,
		 &userInfo.Last_Name, &userInfo.Image); if err != nil {
		return models.UserInfo{}, err
	}
	return userInfo, nil
}