package noteserrors

import "errors"

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrInternal     = errors.New("error occured on the server side")
	ErrTimedOut     = errors.New("timed out")
	ErrNotFound     = errors.New("not found")
)
