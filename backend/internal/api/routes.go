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
	mux.HandleFunc("/api/groups", handler.GetGroups)
	mux.HandleFunc("/api/group", handler.GetGroup)

	// followers
	mux.HandleFunc("/api/followers", handler.HandleGetFollowers) //get
	mux.HandleFunc("/api/followings", handler.HandleGetFollowings) //get
	mux.HandleFunc("/api/follow/requests", handler.HandleGetFollowRequests) //get
	mux.HandleFunc("/api/follow", handler.HandleFollow) //post {"user_id":2}
	mux.HandleFunc("/api/follow/decision", handler.HandleFollowRequest) //post {"follower":2,"status":"accepted"}


	return mux
}
