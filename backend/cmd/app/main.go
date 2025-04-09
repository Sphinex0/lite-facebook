package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"social-network/internal/api"
	"social-network/internal/repository"
	"social-network/pkg/middlewares"
	"social-network/pkg/ratelimiter"
)

func main() {
	db, err := repository.OpenDb()
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			db.Close()
			log.Fatal("Error: ", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		db.Close()
		fmt.Println()
		os.Exit(0)
	}()

	// Set flags to include file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := repository.ApplyMigrations(db); err != nil {
		panic("Migration failed: " + err.Error())
	}

	baseHandler := api.Routes(db)
	rateLimitedHandler := ratelimiter.CreateArticleLimiter.RateMiddleware(baseHandler, 20, 100*time.Millisecond)
	authHandler := middlewares.AuthMiddleware(rateLimitedHandler, db)
	finalHandler := middlewares.ErrorHandler(middlewares.CORS(authHandler))

	server := &http.Server{
		Addr:    ":8080",
		Handler: finalHandler,
	}
	log.Println("http://localhost:8080/")
	err = server.ListenAndServe()
	if err != nil {
		log.Println("Error in starting of server:", err)
		db.Close()
		return
	}
}
