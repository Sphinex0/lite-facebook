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

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler, db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		allowedPath := []string{"/api/login", "/api/signup"}
		Hasallowed := slices.IndexFunc(allowedPath, func(ext string) bool {
			return strings.Contains(r.URL.Path, ext)
		})
		
		cookie, err := r.Cookie("uuid")
		if err != nil && Hasallowed == -1 {
			utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		if cookie != nil {
			uuid, err := uuid.FromString(cookie.Value)
			if err != nil && Hasallowed == -1 {
				utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			if err == nil {
				user, err := repository.GetUserByUuid(db, uuid)
				if err != nil && Hasallowed == -1 {
					utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
					return
				}

				if err == nil {
					ctx := context.WithValue(r.Context(), UserIDKey, user)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
