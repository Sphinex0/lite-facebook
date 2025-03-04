package models

// Conversation represents a conversation between users.
type Conversation struct {
	ID       string `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Type     string `json:"type"`
}
