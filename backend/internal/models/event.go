package models

// Event represents an event within a group.
type Event struct {
	ID          int `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	GroupID     int `json:"group_id"`
	UserID      int `json:"user_id"`
}
