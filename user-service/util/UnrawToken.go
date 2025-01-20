package util

import (
	"strings"

	"github.com/izya4ka/notes-web/user-service/myerrors"
)

func UnrawToken(auth_header string) (string, error) {
	if strings.HasPrefix(auth_header, "Bearer ") {
		return strings.TrimPrefix(auth_header, "Bearer "), nil
	} else {
		return "", myerrors.ErrInvalidToken(0)
	}
}
