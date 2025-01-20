package myerrors

import "fmt"

type ErrNotWithoutSpecSym string

func (err ErrNotWithoutSpecSym) Error() string {
	return fmt.Sprintf("Username '%s' can't contain symbol characters!", string(err))
}

type ErrAlreadyExists string

func (err ErrAlreadyExists) Error() string {
	return fmt.Sprintf("User '%s' already exists!", string(err))
}

type ErrJSON struct {
	Msg string `json:"msg"`
	Error string `json:"error"`
}

type ErrStringLen struct {
	StringName   string
	MustSize     int
	LessOrBigger string
}

func (err ErrStringLen) Error() string {
	return fmt.Sprintf("Length of %s must be %s than %d!", err.StringName, err.LessOrBigger, err.MustSize)
}

type ErrUserNotFound string

func (err ErrUserNotFound) Error() string {
	return fmt.Sprintf("User '%s' not found!", string(err))
}

type ErrMismatchPass string

func (err ErrMismatchPass) Error() string {
	return fmt.Sprintf("Password for user '%s' mismatched!", string(err))
}

type ErrInvalidToken int

func (err ErrInvalidToken) Error() string {
	return "Invalid token!"
}