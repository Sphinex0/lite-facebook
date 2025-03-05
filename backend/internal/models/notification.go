package models

// Notification represents a user notification.
type Notification struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	Type      string `json:"type"`
	InvokerID int `json:"invoker_id"`
	GroupID   int `json:"group_id"`
	EventID   int `json:"event_id"`
}
