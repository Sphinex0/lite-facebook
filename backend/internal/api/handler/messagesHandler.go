package handler

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"social-network/internal/models"
	utils "social-network/pkg"
	"social-network/pkg/middlewares"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[string]*websocket.Conn)
	clientsMu sync.RWMutex
)

func (Handler *Handler) MessagesHandler(upgrader websocket.Upgrader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//
		cookie, err := r.Cookie("session_id")
		if err != nil {
			utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		session := cookie.Value

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("WebSocket Upgrade Error:", err)
			utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		fmt.Println("WebSocket Connected")
		defer conn.Close()

		user, ok := r.Context().Value(middlewares.UserIDKey).(models.User)
		if !ok {
			utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		key := fmt.Sprint(user.ID) + "_" + session
		clientsMu.Lock()
		clients[key] = conn
		clientsMu.Unlock()
		broadcastUserStatus()
		defer removeClient(session)

		for {
			var msg models.WSMessage
			if err := conn.ReadJSON(&msg); err != nil {
				fmt.Println("WebSocket Read Error:", err)
				break
			}
			msg.Message.SenderID = user.ID
			handleMessage(msg, conn, Handler)
		}
	}
}

func handleMessage(msg models.WSMessage, conn *websocket.Conn, Handler *Handler) {
	var err error
	if msg.Kind == "private" {
		switch msg.Type {
		case "new_message":
			msg.Message.Content = strings.TrimSpace(msg.Message.Content)
			if len(msg.Message.Content) == 0 || len(msg.Message.Content) > 500 {
				ErrorMessage(msg.Message.SenderID, conn)
				return
			}

			// if err := msg.Message.StoreMessage(); err != nil {
			// 	fmt.Println("Message Store Error:", err)
			// 	ErrorMessage(msg.SenderID, conn)
			// 	return
			// }

			// distributeMessage(msg)
		case "read":
			// msg.Message.UpdateRead()
		case "conversations":
			msg.Type = "conversations"
			clientsMu.RLock()
			msg.OnlineUsers = getClientIDs()
			clientsMu.RUnlock()
			msg.Conversations, err = Handler.Service.FetchConversations(msg.Message.SenderID)
			if err != nil {
				fmt.Println("GetConversations Error:", err)
				return
			}
			conn.WriteJSON(msg)
		case "typing":
			// distributeMessage(msg)
		}
	} else if msg.Kind == "group" {
	}
}

func broadcastUserStatus() {
}

func removeClient(session string) {
	clientsMu.Lock()
	if conn, ok := clients[session]; ok {
		conn.Close()
		delete(clients, session)
	}
	clientsMu.Unlock()

	broadcastUserStatus()
}

func ParseIdUuid(id_session string) (id string, session string) {
	tab := strings.Split(id_session, "_")
	id, session = tab[0], tab[1]
	return
}

func ErrorMessage(senderID int, conn *websocket.Conn) {
	conn.WriteJSON(models.WSMessage{
		Type: "error",
	})
}

func getClientIDs() []string {
	keys := make([]string, 0, len(clients))
	for key := range clients {
		keys = append(keys, key)
	}
	return keys
}
