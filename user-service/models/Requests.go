package models

// LogPassRequest represents a request to log in with a username and password.
type LogPassRequest struct {
	Username string `json:"username"` // The user's username.
	Password string `json:"password"` // The user's password.
}

// UserChangeCredsRequest represents a request to change a user's credentials.
type UserChangeCredsRequest struct {
	Username    string `json:"username"`     // The user's current username.
	NewUsername string `json:"new_username"` // The user's new username.
	NewPassword string `json:"new_password"` // The user's new password.
}
