package api

import (
	"database/sql"
	"net/http"

	"social-network/internal/api/handler"
)

func Routes(db *sql.DB) *http.ServeMux {
	handler := handler.NewHandler(db)
	mux := http.NewServeMux()

	mux.HandleFunc("/api/Login", handler.Login)
	mux.HandleFunc("/Signup", handler.Signup)

	return mux
}
