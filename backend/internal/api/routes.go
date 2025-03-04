package api

import (
	"database/sql"
	"net/http"
)

func Routes(db *sql.DB) *http.ServeMux {
	//	H := handler.NewHandler(db)
	mux := http.NewServeMux()
	/*
		mux.HandleFunc("/api/Login", H.Login)
		mux.HandleFunc("/Signup", H.Signup)
	*/
	return mux
}
