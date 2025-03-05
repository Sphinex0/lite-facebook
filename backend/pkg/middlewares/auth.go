package middlewares

import (
	"context"
	"database/sql"
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

		cookie, err := r.Cookie("session_id")
		if Hasallowed == -1 {
			if err != nil {
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
