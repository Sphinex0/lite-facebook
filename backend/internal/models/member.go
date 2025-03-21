package models

type Member struct {
	ID             int `json:"id"`
	Member         int `json:"member"`
	ConversationId int `json:"conversation_id"`
	Seen           int `json:"seen"`
}
