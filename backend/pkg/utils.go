package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, statuscode int, Data any) error {
	w.WriteHeader(statuscode)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Data)
	if err != nil {
		return err
	}
	return nil
}

/*
// transform the next set and delete cookie to gorillamux once //

func SetSessionCookie(w http.ResponseWriter, uid string) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  uid,
		Path:   "/",
		MaxAge: 3600,
	})
}

func DeleteSessionCookie(w http.ResponseWriter, uid string) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  uid,
		Path:   "/",
		MaxAge: -1,
	})
}

func Contains(slice []string, str string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == str {
			return true
		}
	}
	return false
}
*/