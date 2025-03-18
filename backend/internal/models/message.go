package models

// Message represents a message in a conversation.
type Message struct {
	ID             int    `json:"id"`
	ConversationID int    `json:"conversation_id"`
	SenderID       int    `json:"sender_id"`
	Content        string `json:"content"`
	Seen           int    `json:"seen"`
	Image          string `json:"image"`
	Reply          *int   `json:"reply"`
	CreatedAt      int    `json:"created_at"`
	ModifiedAt     int    `json:"modified_at"`
}

type WSMessage struct {
	Message       Message             `json:"message"`
	UserInfo      UserInfo            `json:"user_info"`
	Type          string              `json:"type"`
	Typing        bool                `json:"is_typing"`
	Conversations []ConversationsInfo `json:"conversations"`
	OnlineUsers   []int               `json:"online_users"`
}
