package usererrors

import (
	"errors"
	"fmt"
)

// ErrStringLen represents an error related to string length validation.
type ErrStringLen struct {
	StringName    string // The name of the string being validated
	MustSize      int    // The required size of the string
	LessOrGreater string // Indicates if the size must be less or greater
}

// Error formats the error message for string length expectations.
func (err ErrStringLen) Error() string {
	return fmt.Sprintf("length of %s must be %s than %d", err.StringName, err.LessOrGreater, err.MustSize)
}

var (
	ErrInvalidToken      = errors.New("invalid token")
	ErrInternal          = errors.New("error occured on the server side")
	ErrTimedOut          = errors.New("timed out")
	ErrMismatchPass      = errors.New("wrong password")
	ErrUserNotFound      = errors.New("user not found")
	ErrAlreadyExists     = errors.New("user already exists")
	ErrNotWithoutSpecSym = errors.New("username can't contain symbol characters")
)
