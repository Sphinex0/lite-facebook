package models

// Notification represents a user notification.
type Notification struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Type      string `json:"type"`
	InvokerID string `json:"invoker_id"`
	GroupID   string `json:"group_id"`
	EventID   string `json:"event_id"`
}
