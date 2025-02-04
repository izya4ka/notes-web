package usererrors

import (
	"errors"
	"fmt"
)

// ErrNotWithoutSpecSym represents an error when a username contains symbol characters.
type ErrNotWithoutSpecSym string

// Error formats the error message indicating that the username cannot contain symbol characters.
func (err ErrNotWithoutSpecSym) Error() string {
	return fmt.Sprintf("username '%s' can't contain symbol characters", string(err))
}

// ErrAlreadyExists signifies an error that occurs when attempting to create a user that already exists.
type ErrAlreadyExists string

// Error formats the error message indicating that the user already exists.
func (err ErrAlreadyExists) Error() string {
	return fmt.Sprintf("user '%s' already exists", string(err))
}

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

// ErrUserNotFound signifies an error when a user cannot be found.
type ErrUserNotFound string

// Error formats the error message indicating that the user was not found.
func (err ErrUserNotFound) Error() string {
	return fmt.Sprintf("user '%s' not found", string(err))
}

// ErrMismatchPass represents an error when a provided password does not match the expected password for a user.
type ErrMismatchPass string

// Error formats the error message indicating a password mismatch for the specified user.
func (err ErrMismatchPass) Error() string {
	return fmt.Sprintf("password for user '%s' mismatched", string(err))
}

var ErrInvalidToken = errors.New("invalid token")
var ErrInternal = errors.New("error occured on the server side")
