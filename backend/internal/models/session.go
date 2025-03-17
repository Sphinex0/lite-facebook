package models

// Session represents a user session.
type Session struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	SessionID int `json:"session_id"`
	SessionExp int `json:"session_exp"`
}
