package noteserrors

import "errors"

var ErrInvalidToken = errors.New("invalid token")
var ErrInternal = errors.New("error occured on the server side")
var ErrTimedOut = errors.New("timed out")
