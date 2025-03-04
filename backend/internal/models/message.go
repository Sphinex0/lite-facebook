package models

// Message represents a message in a conversation.
type Message struct {
	ID             string `json:"id"`
	ConversationID string `json:"conversation_id"`
	SenderID       string `json:"sender_id"`
	Content        string `json:"content"`
	CreatedAt      string `json:"created_at"`
	Read           string `json:"read"`
}
