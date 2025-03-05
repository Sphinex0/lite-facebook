package models

// Session represents a user session.
type Session struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	SessionExp int `json:"session_exp"`
}
