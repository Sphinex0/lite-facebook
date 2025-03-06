package api

import (
	"database/sql"
	"net/http"

	"social-network/internal/api/handler"
)

func Routes(db *sql.DB) *http.ServeMux {
	handler := handler.NewHandler(db)
	mux := http.NewServeMux()

	// log
	mux.HandleFunc("/api/login", handler.Login)
	mux.HandleFunc("/api/signup", handler.Signup)
	mux.HandleFunc("/api/logout", handler.Logout)

	// profile
	mux.HandleFunc("/api/user", handler.GetUser)
	mux.HandleFunc("/api/user/update", handler.UpdateUser)

	// articls
	mux.HandleFunc("/api/posts", handler.GetArticles)
	mux.HandleFunc("/api/posts/store", handler.AddPost)
	mux.HandleFunc("/api/comments", handler.GetArticles)
	mux.HandleFunc("/api/comments/store", handler.AddComment)
	mux.HandleFunc("/api/reactions/store", handler.GetArticles)

	// group
	mux.HandleFunc("/api/CreateGroup", handler.AddGroup)
	mux.HandleFunc("/api/groups", handler.GetGroups)
	mux.HandleFunc("/api/group/{id}", handler.GetGroup)

	// followers
	mux.HandleFunc("/api/followers", handler.GetFollowers)
	mux.HandleFunc("/api/followings", handler.GetFollowings)

	return mux
}
