package service

import (
	"errors"
	"fmt"
	"html"
	"net/http"
	"net/mail"
	"strings"
	"time"

	"social-network/internal/models"
	utils "social-network/pkg"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (S *Service) LoginUser(User *models.User) (string, error) {
	// check email len
	if ValidateLength(User.Email) {
		return "", errors.New("too long or too short Email")
	}

	// check password len
	if ValidateLength(User.Password) {
		return "", errors.New("too long or too short Password")
	}

	// chefk password and email validity
	id, err := S.Database.CheckMailAndPaswdvalidity(User.Email, User.Password)
	if err != nil {
		return "", err
	}

	User.ID = id
	// generate new uuid
	Uuid := GenerateUuid()

	// Update uuid
	S.Database.AddUuid(Uuid, User.ID)
	return Uuid, nil
}

func (s *Service) RegisterUser(user *models.User) (string, error, int) {
	// First_Name
	if len((*user).First_Name) < 3 || len((*user).First_Name) > 15 {
		return "", errors.New("InvalideFirst_Name"), 0
	}

	// Last_Name
	if len((*user).Last_Name) < 3 || len((*user).Last_Name) > 15 {
		return "", errors.New("InvalideLast_Name"), 0
	}

	// Password
	if len((*user).Password) < 6 || len((*user).Password) > 30 {
		return "", errors.New("InvalidPassword"), 0
	}

	// email
	(*user).Email = strings.ToLower((*user).Email)
	if !EmailChecker((*user).Email) {
		return "", errors.New("InvalidEmail"), 0
	}
	if len((*user).Email) > 50 {
		return "", errors.New("LongEmail"), 0
	}

	// dob
	err := validateDOB(int64((*user).DateBirth))
	if err != nil {
		return "", err, 0
	}

	// Nickname
	if (*user).Nickname != "" {
		if len(strings.TrimSpace((*user).Nickname)) < 3 || len(strings.TrimSpace((*user).Nickname)) > 15 {
			return "", errors.New("InvalidUsername"), 0
		}
	}

	// AboutMe
	if (*user).AboutMe != "" {
		if len(strings.TrimSpace((*user).AboutMe)) > 50 {
			return "", errors.New("InvalidUsername"), 0
		}
	}

	// username or email existance
	if s.Database.CheckIfUserExists((*user).Email) {
		return "", errors.New("UserAlreadyExist"), 0
	}

	// Encrypt Pass
	(*user).Password, err = EncyptPassword((*user).Password)
	if err != nil {
		return "", err, 0
	}

	// Fix username html
	(*user).Nickname = html.EscapeString((*user).Nickname)

	// Generate Uuid
	Uuid := GenerateUuid()
	// Insert the user
	id, err := s.Database.InsertUser(*user, Uuid)
	if err != nil {
		return "", err, 0
	}

	return Uuid, nil, id
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
	err := S.Database.DeleteCookieFromdb(uuid)
	if err != nil {
		return errors.New("error while deleting the cookie")
	}
	utils.DeleteSessionCookie(w, uuid)
	return nil
}

func (S *Service) Extractuser(r *http.Request) models.User {
	user := models.User{
		Email:      r.FormValue("email"),
		Password:   r.FormValue("password"),
		First_Name: r.FormValue("firstName"),
		Last_Name:  r.FormValue("lastName"),
		Nickname:   r.FormValue("nickname"),
		AboutMe:    r.FormValue("aboutMe"),
	}
	return user
}

func validateDOB(dobStr int64) error {
	dob := time.UnixMilli(dobStr)
	age := time.Since(dob)
	year := 365 * 24 * time.Hour
	minAge := 18 * year
	maxAge := 100 * year
	if age < minAge {
		return fmt.Errorf("you must be at least 18 years old")
	}
	if age > maxAge {
		return fmt.Errorf("invalid age, please check the date of birth")
	}
	return nil
}
