package models

// Event represents an event within a group.
type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Name        string `json:"name"`
	GroupID     string `json:"group_id"`
	UserID      string `json:"user_id"`
}
