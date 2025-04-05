package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"

	"social-network/internal/models"

	"github.com/gofrs/uuid/v5"
)

type contextKey string

const UserIDKey contextKey = "user"

const UserCookie contextKey = "cookie"

func WriteJson(w http.ResponseWriter, statuscode int, Data any) error {
	w.WriteHeader(statuscode)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Data)
	if err != nil {
		return err
	}
	return nil
}

func ParseBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

// GetScanFields returns a slice of pointers to struct fields for scanning SQL results.
// Pass a pointer to the struct.
// Example: GetScanFields(&user) => []*interface{}{&user.ID, &user.FirstName, &user.LastName, &user.Nickname, &user.Image}
func GetScanFields(s interface{}) []interface{} {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		log.Fatal("Input must be a pointer to a struct")
	}
	val = val.Elem() // like &user => user

	fields := make([]interface{}, val.NumField()) // like user => user.ID, user.FirstName, ...
	for i := 0; i < val.NumField(); i++ {
		fields[i] = val.Field(i).Addr().Interface()
	}
	return fields
}

// GetExecFields returns a slice of struct field values, excluding specified fields.
// Example: GetExecFields(user, "ID", "CreatedAt") => []interface{}{user.FirstName, user.LastName, user.Nickname, user.Image}
func GetExecFields(s interface{}, excludeFields ...string) []interface{} {
	val := reflect.ValueOf(s) // like user => user.FirstName, user.LastName, ...
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		log.Fatal("Input must be a struct or pointer to a struct")
	}

	// Convert excluded fields to a map for fast lookup
	excluded := make(map[string]bool)
	for _, field := range excludeFields {
		excluded[field] = true
	}

	var fields []interface{}
	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		if !excluded[fieldName] {
			fields = append(fields, val.Field(i).Interface())
		}
	}
	return fields
}

func SetSessionCookie(w http.ResponseWriter, uuid string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    uuid,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   31536000,
	})
}

func DeleteSessionCookie(w http.ResponseWriter, uuid string) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  uuid,
		Path:   "/",
		MaxAge: -1,
	})
}

func Length(a, b int, e string) bool {
	return len(e) < a || len(e) > b
}

func Placeholders(n int) string {
	return strings.Repeat("?,", n)[:2*n-1]
}

// type ContextKey string

// const UserIDKey ContextKey = "user"

func GetUserFromContext(ctx context.Context) (user models.UserInfo, uuidCookie string, ok bool) {
	user, ok = ctx.Value(UserIDKey).(models.UserInfo)
	if !ok {
		return
	}
	uuidCookie, ok = ctx.Value(UserCookie).(string)
	return
}

func StoreThePic(UploadDir string, file multipart.File, handler *multipart.FileHeader) (string, error) {
	if _, err := os.Stat(UploadDir); os.IsNotExist(err) {
		fmt.Println(err)
		er := os.Mkdir(UploadDir, os.ModePerm)
		fmt.Println(er)
	} else {
		fmt.Println(err)
	}

	randomstr := GenerateUuid()
	fmt.Println(randomstr + handler.Filename)

	extensions := []string{".png", ".jpg", ".jpeg", ".gif"}
	extIndex := slices.IndexFunc(extensions, func(ext string) bool {
		return strings.HasSuffix(strings.ToLower(handler.Filename), ext)
	})
	if extIndex == -1 {
		return "" , fmt.Errorf("err")
	}

	filePath := filepath.Join(UploadDir, randomstr+extensions[extIndex])
	dst, err := os.Create(filePath)
	if err != nil {
		fmt.Println("err1",err)
		return "", errors.New("could not save file")
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println("err2",err)
		return "", errors.New("failed to save file")
	}

	return randomstr+extensions[extIndex], nil
}

func GenerateUuid() string {
	return uuid.Must(uuid.NewV4()).String()
}
