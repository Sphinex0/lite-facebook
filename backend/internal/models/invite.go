package models

// Invite represents an invitation to a group.
type Invite struct {
	ID       int `json:"id"`
	GroupID  int `json:"group_id"`
	Sender   int `json:"sender"`
	Receiver int `json:"receiver"`
	Status   string `json:"status"`
}
