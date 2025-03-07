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
	mux.HandleFunc("/api/posts", handler.HandelGetArticles)
	mux.HandleFunc("/api/comments", handler.HandelGetArticles)
	mux.HandleFunc("/api/articles/store", handler.HandelCreateArticle)
	mux.HandleFunc("/api/reactions/store", handler.HandelCreateReaction)

	// group
	mux.HandleFunc("/api/groups/store", handler.AddGroup)
	mux.HandleFunc("/api/groups", handler.GetGroups)
	mux.HandleFunc("/api/group/{id}", handler.GetGroup)

	// Invites
	mux.HandleFunc("/api/invate/store//{IdGroup}/{IdReciver}", handler.AddInviteByReceiver)
	mux.HandleFunc("/api/addInvitebySender/{IdGroup}/{IdReciver}", handler.AddInviteBySender)
	mux.HandleFunc("/api/Invite/{id}", handler.GetGroup)


	// followers
	mux.HandleFunc("/api/followers", handler.GetFollowers)
	mux.HandleFunc("/api/followings", handler.GetFollowings)

	return mux
}
