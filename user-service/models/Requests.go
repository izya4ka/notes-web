package models

// LogPassRequest represents a request to log in with a username and password.
type LogPassRequest struct {
	Username string `json:"username"` // The user's username.
	Password string `json:"password"` // The user's password.
}
