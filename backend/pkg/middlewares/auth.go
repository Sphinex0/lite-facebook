package middlewares

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"slices"
	"strings"

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
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, withCredentials")
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

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler, db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedPath := []string{"/api/login", "/api/signup"}
		Hasallowed := slices.IndexFunc(allowedPath, func(ext string) bool {
			return strings.Contains(r.URL.Path, ext)
		})

		cookie, err := r.Cookie("session_id")
		fmt.Println(r.Cookies())
		if Hasallowed == -1 {
			if err != nil {
				fmt.Println(err)
				utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
				return
			}
			uuid, err := uuid.FromString(cookie.Value)
			if err != nil {
				utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			user, err := repository.GetUserByUuid(db, uuid)
			if err != nil {
				utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
				return
			}
			ctx := context.WithValue(r.Context(), UserIDKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {
			if err == nil {
				uuid, err := uuid.FromString(cookie.Value)
				if err == nil {
					_, err := repository.GetUserByUuid(db, uuid)
					if err == nil {
						utils.WriteJson(w, http.StatusForbidden, "StatusForbidden")
						return
					}
				}
			}
			next.ServeHTTP(w, r)
		}
	})
}
