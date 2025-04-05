package models

// Event represents an event within a group.
type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Day         string `json:"day"`
	GroupID     int    `json:"group_id"`
	UserID      int    `json:"user_id"`
}

type EventPardse struct {
	Event Event `json="event"`
	Going bool  `json="going"`
}
