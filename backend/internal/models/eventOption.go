package models

// EventOption represents a user's response to an event.
type EventOption struct {
	ID      int `json:"id"`
	Going   int `json:"going"`
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}
