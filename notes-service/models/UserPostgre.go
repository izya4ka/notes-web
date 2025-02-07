package models

// UserPostgres represents a user in the PostgreSQL database.
// It includes fields for storing user information and
// establishes a one-to-many relationship with the Note model.
// The ID field serves as the primary key, while CreatedAt and
// UpdatedAt fields are automatically managed by Gorm for
// tracking record creation and modification times.
type UserPostgres struct {
	ID        uint   `gorm:"primaryKey" json:"id"`                                 // Unique identifier for the user.
	Username  string `json:"username" gorm:"uniqueIndex"`                          // The user's chosen username.
	Password  string `json:"password"`                                             // The user's password, should be securely hashed.
	Token     string `json:"token" gorm:"uniqueIndex"`                             // Current user's auth token
	Notes     []Note `json:"notes" gorm:"foreignKey:Username;references:Username"` // List of notes associated with the user.
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`               // Timestamp of when the user was created.
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`               // Timestamp of the last update to the user's record.
}
