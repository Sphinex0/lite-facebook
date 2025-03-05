package models

// Conversation represents a conversation between users.
type Conversation struct {
	ID       int `json:"id"`
	Entitie_one   int `json:" entitie_one"`
	Entitie_two int `json:" entitie_two"`
	Type     string `json:"type"`
}
