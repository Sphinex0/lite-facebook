package models

// Message represents a message in a conversation.
type Message struct {
	ID             int `json:"id"`
	ConversationID int `json:"conversation_id"`
	SenderID       int `json:"sender_id"`
	Content        string `json:"content"`
	CreatedAt      int `json:"created_at"`
	Read           int `json:"read"`
}
