package models

// Invite represents an invitation to a group.
type Invite struct {
	ID       string `json:"id"`
	GroupID  string `json:"group_id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
}
