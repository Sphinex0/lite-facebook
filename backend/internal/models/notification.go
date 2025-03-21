package models

import "database/sql"

// Notification represents a user notification.
type Notification struct {
	ID          int            `json:"id"`
	UserID      int            `json:"user_id"`
	Type        string         `json:"type"`
	InvokerID   int            `json:"invoker_id"`
	InvokerName sql.NullString `json:"invoker_name"`
	GroupID     int            `json:"group_id"`
	GroupTitle  sql.NullString `json:"group_title"`
	EventID     int            `json:"event_id"`
	EventName   string         `json:"event_name"`
	Seen        bool           `json:"seen"`
}
