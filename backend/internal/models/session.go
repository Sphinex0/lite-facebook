package models

// Session represents a user session.
type Session struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	SessionExp string `json:"session_exp"`
}
