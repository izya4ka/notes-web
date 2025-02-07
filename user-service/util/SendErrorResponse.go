package util

import (
	"errors"
	"net/http"
	"time"

	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/usererrors"
	"github.com/labstack/echo/v4"
)

func SendErrorResponse(c echo.Context, err error) error {

	var error_response models.Error

	switch err {
	case usererrors.ErrInvalidToken:
		error_response.Code = http.StatusUnauthorized
	case usererrors.ErrAlreadyExists:
		error_response.Code = http.StatusConflict
	case usererrors.ErrInternal:
		error_response.Code = http.StatusInternalServerError
	case usererrors.ErrMismatchPass:
		error_response.Code = http.StatusUnauthorized
	case usererrors.ErrTimedOut:
		error_response.Code = http.StatusRequestTimeout
	case usererrors.ErrNotWithoutSpecSym:
		error_response.Code = http.StatusBadRequest
	case usererrors.ErrUserNotFound:
		error_response.Code = http.StatusNotFound
	default:
		if errors.As(err, &usererrors.ErrStringLen{}) {
			error_response.Code = http.StatusBadRequest
		}
	}

	error_response.Error = http.StatusText(error_response.Code)
	error_response.Message = err.Error()
	error_response.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	error_response.Path = c.Path()

	return c.JSON(error_response.Code, error_response)
}
