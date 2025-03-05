package models

// Follower represents a follower relationship.
type Follower struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	Follower int `json:"follower"`
	Status   string `json:"status"`
}
