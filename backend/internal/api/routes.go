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
	mux.HandleFunc("/api/posts", handler.HandelGetPosts)
	mux.HandleFunc("/api/comments", handler.HandelGetComments)
	mux.HandleFunc("/api/articles/store", handler.HandelCreateArticle)
	mux.HandleFunc("/api/reactions/store", handler.HandelCreateReaction)

	// group
	mux.HandleFunc("/api/groups/store", handler.AddGroup)
	mux.HandleFunc("/api/groups", handler.GetGroups)
	mux.HandleFunc("/api/group/{id}", handler.GetGroup)

	// Invites
	mux.HandleFunc("/api/invate/store", handler.AddInviteByReceiver)
	mux.HandleFunc("/api/invite/decision", handler.HandleInviteRequest)


	// followers
	mux.HandleFunc("/api/followers", handler.GetFollowers)
	mux.HandleFunc("/api/followings", handler.GetFollowings)
	mux.HandleFunc("/api/follow", handler.HandleFollow)
	mux.HandleFunc("/api/follow/decision", handler.HandleFollowRequest)


	return mux
}
