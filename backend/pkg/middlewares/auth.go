package middlewares

import (
	"context"
	"database/sql"
	"net/http"
	"slices"

	"social-network/internal/repository"
	utils "social-network/pkg"

	"github.com/gofrs/uuid/v5"
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

type contextKey string

const UserIDKey contextKey = "user"

func AuthMiddleware(next http.Handler, db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Consider specifying allowed origins
		allowedPaths := []string{"/api/login", "/api/signup"}
		isAllowedPath := slices.Contains(allowedPaths, r.URL.Path)

		cookie, err := r.Cookie("session_id")
		if !isAllowedPath {
			// Path requires authentication
			if err != nil {
				utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
				return
			}
			sessionUUID, err := uuid.FromString(cookie.Value)
			if err != nil {
				utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			user, err := repository.GetUserByUuid(db, sessionUUID)
			if err != nil {
				utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
				return
			}
			ctx := context.WithValue(r.Context(), UserIDKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			// Path is allowed (login/signup), restrict if authenticated
			if cookie != nil {
				sessionUUID, err := uuid.FromString(cookie.Value)
				if err == nil {
					_, err := repository.GetUserByUuid(db, sessionUUID)
					if err == nil {
						utils.WriteJson(w, http.StatusForbidden, "Forbidden")
						return
					}
				}
			}
			next.ServeHTTP(w, r)
		}
	})
}
