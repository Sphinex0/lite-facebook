package models

// Article represents an article entity, corrected from "Articls".
type Article struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	Content    string `json:"content"`
	Privacy    string `json:"privacy"`
	CreatedAt  int `json:"created_at"` 
	ModifiedAt int `json:"modified_at"`
	Image      string `json:"image"`
	Parent     int `json:"parent"`
	GroupID    int `json:"group_id"`
}
