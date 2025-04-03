package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"social-network/internal/api"
	"social-network/internal/repository"
	"social-network/pkg/middlewares"
	"social-network/pkg/ratelimiter"
)

func main() {
	db, err := repository.OpenDb()
	if err != nil {
		return
	}

	// Set flags to include file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := repository.ApplyMigrations(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	baseHandler := api.Routes(db)

	rateLimitedHandler := ratelimiter.CreateArticleLimiter.RateMiddleware(baseHandler, 20, 100*time.Millisecond)

	authHandler := middlewares.AuthMiddleware(rateLimitedHandler, db)

	finalHandler := middlewares.CORS(authHandler)
	
	server := &http.Server{
		Addr:    ":8080",
		Handler: finalHandler,
	}
	fmt.Println("http://localhost:8080/")
	err = server.ListenAndServe()
	if err != nil {
		log.Println("Error in starting of server:", err)
		return
	}
}
