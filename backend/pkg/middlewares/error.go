package middlewares

import (
	"log"
	"net/http"

	utils "social-network/pkg"
)

// Centralized Error Handling Middleware
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		
		next.ServeHTTP(w, r)
	})
}
