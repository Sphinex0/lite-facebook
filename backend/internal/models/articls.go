package models

// Article represents an article entity, corrected from "Articls".
type Article struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Content    string `json:"content"`
	Privacy    string `json:"privacy"`
	CreatedAt  string `json:"created_at"` 
	ModifiedAt string `json:"modified_at"`
	Image      string `json:"image"`
	Parent     string `json:"parent"`
	GroupID    string `json:"group_id"`
}
