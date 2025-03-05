package models

// EventOption represents a user's response to an event.
type EventOption struct {
	ID      int `json:"id"`
	Going   string `json:"going"`
	UserID  string `json:"user_id"`
	EventID string `json:"event_id"`
}
