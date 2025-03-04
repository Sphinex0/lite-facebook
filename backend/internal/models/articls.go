package models

// Article represents an article entity, corrected from "Articls".
type Article struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"` // Renamed to UserID for Go naming conventions
	Content    string `json:"content"`
	Privacy    string `json:"privacy"`
	CreatedAt  string `json:"created_at"`  // Renamed to CreatedAt for consistency
	ModifiedAt string `json:"modified_at"` // Renamed to ModifiedAt
	Image      string `json:"image"`
	Parent     string `json:"parent"`
	GroupID    string `json:"group_id"` // Renamed to GroupID
}
