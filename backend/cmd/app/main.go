package main

import (
	"fmt"
	"log"
	"net/http"

	"social-network/internal/api"
	"social-network/internal/repository"
	"social-network/pkg/middlewares"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := repository.OpenDb()
	if err != nil {
		fmt.Println("Error in opening of database:", err)
		return
	}

	if err := repository.ApplyMigrations(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: middlewares.CORS(api.Routes(db)),
	}

	fmt.Println("http://localhost:8080/")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Error in starting of server:", err)
		return
	}
}
