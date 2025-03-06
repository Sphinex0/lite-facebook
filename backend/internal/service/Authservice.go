package service

import (
	"errors"
	"html"
	"net/http"
	"net/mail"
	"social-network/internal/models"
	utils "social-network/pkg"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (S *Service) LoginUser(User *models.User) error {
	// check email len
	if ValidateLength(User.Email) {
		return errors.New("too long or too short Email")
	}

	// check password len
	if ValidateLength(User.Password) {
		return errors.New("too long or too short Password")
	}

	// chefk password and email validity
	usrId, err := S.Database.CheckMailAndPaswdvalidity(User.Email, User.Password)
	if err != nil {
		return err
	}
	// generate new uuid
	(*User).Uuid = GenerateUuid()

	// Update uuid
	S.Database.UpdateUuid((*User).Uuid, usrId)
	return nil
}

func (s *Service) RegisterUser(user *models.User) error {
	// Username
	if len((*user).Nickname) < 3 || len((*user).Nickname) > 15 {
		return errors.New("InvalidUsername")
	}

	// Age
	if !CheckAgeValidation((*user).Age) {
		return errors.New("InvalideAge")
	}

	// First_Name
	if len((*user).First_Name) < 3 || len((*user).First_Name) > 15 {
		return errors.New("InvalideFirst_Name")
	}

	// Last_Name
	if len((*user).Last_Name) < 3 || len((*user).Last_Name) > 15 {
		return errors.New("InvalideLast_Name")
	}

	// Password
	if len((*user).Password) < 6 || len((*user).Password) > 30 {
		return errors.New("InvalidPassword")
	}

	// email
	(*user).Email = strings.ToLower((*user).Email)
	if !EmailChecker((*user).Email) {
		return errors.New("InvalidEmail")
	}
	if len((*user).Email) > 50 {
		return errors.New("LongEmail")
	}

	// username or email existance
	if s.Database.CheckIfUserExists((*user).Email) {
		return errors.New("UserAlreadyExist")
	}

	// Generate Uuid
	(*user).Uuid = GenerateUuid()

	// Encrypt Pass
	var err error
	(*user).Password, err = EncyptPassword((*user).Password)
	if err != nil {
		return err
	}

	// Fix username html
	(*user).Nickname = html.EscapeString((*user).Nickname)

	// Insert the user
	return s.Database.InsertUser(*user)
}

func ValidateLength(data string) bool {
	if len(data) <= 3 || len(data) >= 32 {
		return true
	}
	return false
}

func GenerateUuid() string {
	return uuid.Must(uuid.NewV4()).String()
}

func CheckAgeValidation(age string) bool {
	Age, err := strconv.Atoi(age)
	if err != nil {
		return false
	}
	if Age > 1000 || Age < 0 {
		return false
	}

	return true
}

func (s *Service) GetInfoData(userUID string) (bool, error) {
	// Get username and user id
	id, _ := s.Database.GetUser(userUID)
	// if id = 0 that means the user doesn't exist
	authorized := id != 0

	return authorized, nil
}

func EmailChecker(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func CheckPasswordValidity(hashedPass, entredPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(entredPass))
	return err == nil
}

func EncyptPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func (S *Service) DeleteSessionCookie(w http.ResponseWriter, uuid string) error {
	err := S.Database.DeleteCookieFromdb(uuid); if err != nil {
		return errors.New("error while deleting the cookie")
	}
	utils.DeleteSessionCookie(w, uuid)
	return nil
}
