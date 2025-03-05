package models

// Conversation represents a conversation between users.
type Conversation struct {
	ID       int `json:"id"`
	Sender   int `json:"sender"`
	Receiver int `json:"receiver"`
	Type     string `json:"type"`
}
