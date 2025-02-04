package util

import (
	"strings"

	"github.com/izya4ka/notes-web/user-service/usererrors"
)

// UnrawToken extracts the token string from an authorization header.
// It checks if the header begins with "Bearer " and, if so, returns the
// token by trimming this prefix. If the header does not have the correct
// format, it returns an empty string and an error indicating that the
// token is invalid.
func UnrawToken(auth_header string) (string, error) {
	if strings.HasPrefix(auth_header, "Bearer ") {
		return strings.TrimPrefix(auth_header, "Bearer "), nil
	} else {
		return "", usererrors.ErrInvalidToken // Return an error for invalid token format
	}
}
