package handler

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
	"sync"

	"social-network/internal/models"
	utils "social-network/pkg"
	"social-network/pkg/middlewares"

	"github.com/gorilla/websocket"
)

var (
	userConnections         = make(map[int][]*websocket.Conn)
	conversationSubscribers = make(map[int][]int)
	connMu                  sync.RWMutex
	subMu                   sync.RWMutex
)

func (Handler *Handler) MessagesHandler(upgrader websocket.Upgrader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		connMu.Lock()
		userConnections[user.ID] = append(userConnections[user.ID], conn)
		connMu.Unlock()

		// broadcastUserStatus()
		defer removeConnection(user.ID, conn)

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

func removeConnection(userID int, conn *websocket.Conn) {
	connMu.Lock()
	defer connMu.Unlock()
	if conns, ok := userConnections[userID]; ok {
		for i, c := range conns {
			if c == conn {
				userConnections[userID] = append(conns[:i], conns[i+1:]...)
				break
			}
		}
		if len(userConnections[userID]) == 0 {
			delete(userConnections, userID)
			fmt.Println("Delete user")
			// broadcastUserStatus()
		}
	}
}

func handleMessage(msg models.WSMessage, conn *websocket.Conn, Handler *Handler) {
	var err error
	switch msg.Type {
	case "new_message":
		subMu.RLock()
		subscribers, ok := conversationSubscribers[msg.Message.ConversationID]
		subMu.RUnlock()
		if !ok || !slices.Contains(subscribers, msg.Message.SenderID) {
			return
		}

		msg.Message.Content = strings.TrimSpace(msg.Message.Content)
		if len(msg.Message.Content) == 0 || len(msg.Message.Content) > 500 {
			ErrorMessage(msg.Message.SenderID, conn)
			return
		}

		if err := Handler.Service.CreateMessage(&msg.Message); err != nil {
			fmt.Println("Message Store Error:", err)
			ErrorMessage(msg.Message.SenderID, conn)
			return
		}

		distributeMessage(msg, subscribers)
	case "read":
		// msg.Message.UpdateRead()
	case "conversations":
		msg.Type = "conversations"
		msg.Conversations, err = Handler.Service.FetchConversations(msg.Message.SenderID)
		if err != nil {
			fmt.Println("GetConversations Error:", err)
			return
		}
		msg.OnlineUsers = getClientIDs(msg.Conversations)
		conn.WriteJSON(msg)
	case "typing":
		// distributeMessage(msg)
	}
}

func broadcastUserStatus() {
}

func ErrorMessage(senderID int, conn *websocket.Conn) {
	conn.WriteJSON(models.WSMessage{
		Type: "error",
	})
}

// distributeMessage sends the message to all conversation participants
func distributeMessage(msg models.WSMessage, subscribers []int) {
	connMu.RLock()
	defer connMu.RUnlock()
	for _, userID := range subscribers {
		if conns, ok := userConnections[userID]; ok {
			for _, conn := range conns {
				if err := conn.WriteJSON(msg); err != nil {
					fmt.Println("Write Error:", err)
				}
			}
		}
	}
}

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func getClientIDs(cnvs []models.ConversationsInfo) (Ids []int) {
	connMu.RLock()
	defer connMu.RUnlock()
	for _, cnv := range cnvs {
		if cnv.Conversation.Type == "private" {
			if _, ok := userConnections[cnv.UserInfo.ID]; ok {
				Ids = append(Ids, cnv.UserInfo.ID)
			}
		}
	}
	return
}
