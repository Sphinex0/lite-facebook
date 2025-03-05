package models

// Group represents a group entity.
type Group struct {
	ID          int `json:"id"`
	Creator     int `json:"creator"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
