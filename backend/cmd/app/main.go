package main

import (
	"fmt"
	"net/http"
	"social-network/internal/api"
	"social-network/internal/repository"
)

func main() {

	db, err := repository.OpenDb()
	if err != nil {
		fmt.Println("Error in opening of database:", err)
		return
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: api.Routes(db),
	}

	fmt.Println("http://localhost:8080/")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Error in starting of server:", err)
		return
	}
}
