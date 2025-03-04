package models

// Follower represents a follower relationship.
type Follower struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Follower string `json:"follower"`
	Status   string `json:"status"`
}
