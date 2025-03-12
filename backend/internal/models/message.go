package models

// Message represents a message in a conversation.
type Message struct {
	ID             int    `json:"id"`
	ConversationID int    `json:"conversation_id"`
	SenderID       int    `json:"sender_id"`
	Content        string `json:"content"`
	Seen           int    `json:"seen"`
	CreatedAt      int    `json:"created_at"`
}

type WSMessage struct {
	Message       Message             `json:"message"`
	Type          string              `json:"type"`
	Typing        bool                `json:"is_typing"`
	Conversations []ConversationsInfo `json:"conversations"`
	OnlineUsers   []int            `json:"online_users"`
}
