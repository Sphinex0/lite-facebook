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

	// Handler: middlewares.CORS(ratelimiter.CreateArticleLimiter.RateMiddleware(middlewares.AuthMiddleware(api.Routes(db), db), 20, 2*time.Second)),
	// auth := middlewares.AuthMiddleware(api.Routes(db), db)
	// handel := ratelimiter.CreateArticleLimiter.RateMiddleware(auth, 20, 2*time.Second)
	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: middlewares.CORS(handel),
	// }

	// Create the base handler
	baseHandler := api.Routes(db)

	// Apply middlewares in the correct order
	// 1. Rate Limiter
	rateLimitedHandler := ratelimiter.CreateArticleLimiter.RateMiddleware(baseHandler, 20, 100*time.Millisecond)

	// 2. Authentication (after rate limiting)
	authHandler := middlewares.AuthMiddleware(rateLimitedHandler, db)

	// 3. CORS (wraps everything last)
	finalHandler := middlewares.CORS(authHandler)

	// Set up and start the server
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
