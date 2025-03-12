package gateerrors

import "errors"

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrInternal     = errors.New("error occured on the server side")
	ErrTimedOut     = errors.New("timed out")
	ErrNotFound     = errors.New("not found")
)

type Error struct {
	Code      int    `json:"code"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
	Timestamp string `json:"timestamp"`
}
