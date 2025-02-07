package models

type Error struct {
	Code      int    `json:"code"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
	Timestamp string `json:"timestamp"`
}

type BaseNoteRequest struct {
	Title       string `json:"title"`       // Title of the note
	Description string `json:"description"` // Detailed description of the note
}

type GetNotesRequest struct {
	Amount int `json:"amount"`
}
