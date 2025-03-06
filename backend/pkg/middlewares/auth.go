package middlewares

import (
	"net/http"
)

// CORS middleware to handle cross-origin requests
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// List of allowed origins (you can modify it based on your needs)
		allowedOrigins := []string{
			"http://localhost:3000",
		}

		// Get the `Origin` header from the request
		origin := r.Header.Get("Origin")

		// Check if the origin is in the allowed origins list
		allowOrigin := "*"
		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == origin {
				allowOrigin = allowedOrigin
				break
			}
		}

		// Set the CORS headers
		w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// If it's a pre-flight request (OPTIONS method), end the request here
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}