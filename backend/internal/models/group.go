package models

// Group represents a group entity.
type Group struct {
	ID          string `json:"id"`
	Creator     string `json:"creator"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
