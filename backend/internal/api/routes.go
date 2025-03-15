package api

import (
	"database/sql"
	"net/http"

	"social-network/internal/api/handler"

	"github.com/gorilla/websocket"
)

func Routes(db *sql.DB) *http.ServeMux {
	handler := handler.NewHandler(db)
	mux := http.NewServeMux()

	// log
	mux.HandleFunc("/api/login", handler.Login)
	mux.HandleFunc("/api/signup", handler.Signup)
	mux.HandleFunc("/api/logout", handler.Logout)
	mux.HandleFunc("/api/checkuser", handler.CheckUserValidity)

	// profile
	mux.HandleFunc("/api/user", handler.GetUser)
	mux.HandleFunc("/api/user/update", handler.UpdateUser)

	// articls
	mux.HandleFunc("/api/posts", handler.HandelGetPosts)                 // post {"before":184525547}
	mux.HandleFunc("/api/comments", handler.HandelGetComments)           // post {"before":184525547, "parent":4}
	mux.HandleFunc("/api/articles/store", handler.HandelCreateArticle)   // post form {"content":"Hello world","privacy":"public" ,"image":file} // or the same but add {"group_id":5} // or the same but add {"parent":5}
	mux.HandleFunc("/api/reactions/store", handler.HandelCreateReaction) // post {"like":1|-1, "article_id":4}
	mux.HandleFunc("/api/group/posts", handler.HandelGetPostsByGroup)    // post {"before":184525547,"group_id":1}

	// group
	mux.HandleFunc("/api/groups/store", handler.AddGroup)
	mux.HandleFunc("/api/groups", handler.GetGroups)
	mux.HandleFunc("/api/group", handler.GetGroup)

	// Invites
	mux.HandleFunc("/api/invite/store", handler.AddInvite)
	mux.HandleFunc("/api/invite/decision", handler.HandleInviteRequest)
	mux.HandleFunc("/api/invites", handler.GetInvites)
	mux.HandleFunc("/api/invites/members", handler.GetMembers)

	// Events
	mux.HandleFunc("/api/Event/store", handler.AddEvent)
	mux.HandleFunc("/api/Events", handler.GetEvents)
	mux.HandleFunc("/api/Event", handler.GetEvent)

	// Events_options

	mux.HandleFunc("/api/Event/options/store", handler.OptionEvent)
	mux.HandleFunc("/api/Event/options", handler.GetEventOption)

	// followers
	mux.HandleFunc("/api/followers", handler.HandleGetFollowers)            // get
	mux.HandleFunc("/api/followings", handler.HandleGetFollowings)          // get
	mux.HandleFunc("/api/follow/requests", handler.HandleGetFollowRequests) // get
	mux.HandleFunc("/api/follow", handler.HandleFollow)                     // post {"user_id":2}
	mux.HandleFunc("/api/follow/decision", handler.HandleFollowRequest)     // post {"follower":2,"status":"accepted"}

	// websocket

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	mux.HandleFunc("/ws", handler.MessagesHandler(upgrader))
	// notification
	mux.HandleFunc("/api/GetNotification", handler.HandleGetNotification) //get


	return mux
}
