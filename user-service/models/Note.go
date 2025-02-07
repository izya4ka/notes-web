package models

// Note represents a user-generated note in the system.
// It contains metadata such as the note's ID, associated user ID, title, and description.
// Timestamps are also included to track creation and updates.
//
// ID is the primary key for the note, ensuring uniqueness.
// UserID associates the note with a specific user in the application.
// Title provides a brief identification of the note's content.
// Description contains the full details of the note's content.
// CreatedAt records the timestamp of when the note was created.
// UpdatedAt records the timestamp of when the note was last updated.
type Note struct {
	ID          uint   `gorm:"primaryKey" json:"id"`                   // Unique identifier for the note
	Username    string `json:"username"`                               // Identifier of the user who created the note
	Title       string `json:"title"`                                  // Title of the note
	Description string `json:"description"`                            // Detailed description of the note
	CreatedAt   int64  `gorm:"autoCreateTime:milli" json:"created_at"` // Timestamp of note creation
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli" json:"updated_at"` // Timestamp of the last update to the note
}
