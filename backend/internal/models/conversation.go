package models

// Conversation represents a conversation between users.
type Conversation struct {
	ID                int    `json:"id"`
	Entitie_one       int    `json:"entitie_one"`
	Entitie_two_user  *int   `json:"entitie_two_user"`
	Entitie_two_group *int   `json:"entitie_two_group"`
	Type              string `json:"type"`
	CreatedAt         int    `json:"created_at"`
	ModifiedAt        int    `json:"modified_at"`
}

type ConversationsInfo struct {
	Conversation Conversation `json:"conversation"`
	UserInfo     UserInfo     `json:"user_info"`
	Group        Group        `json:"group"`
	LastMessage  string       `json:"last_message"`
	Seen         int          `json:"seen"`
}
