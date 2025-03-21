package models

// Notification represents a user notification.
type Notification struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Type        string `json:"type"`
	InvokerID   int    `json:"invoker_id"`
	InvokerName string `json:"invoker_name"`
	GroupID     int    `json:"group_id"`
	GroupTitle  string `json:"group_title"`
	EventID     int    `json:"event_id"`
	EventName   string `json:"event_name"`
	Seen        bool   `json:"seen"`
}
