package util

import (
	"net/http"
	"time"

	"github.com/izya4ka/notes-web/gateway/gateerrors"
	"github.com/labstack/echo/v4"
)

func SendErrorResponse(c echo.Context, err error) error {

	var error_response gateerrors.Error

	error_response.Message = err.Error()

	switch err {
	case gateerrors.ErrInvalidToken:
		error_response.Code = http.StatusUnauthorized
	case gateerrors.ErrInternal:
		error_response.Code = http.StatusInternalServerError
	case gateerrors.ErrTimedOut:
		error_response.Code = http.StatusRequestTimeout
	case gateerrors.ErrNotFound:
		error_response.Code = http.StatusNotFound
	default:
		error_response.Code = http.StatusInternalServerError
		LogErrorf("%s %s %d -> Error: %s", c.Request().Method, c.Path(), error_response.Code, err)
		error_response.Message = "Undocumented Error"
	}

	error_response.Error = http.StatusText(error_response.Code)

	error_response.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	error_response.Path = c.Request().URL.Path

	return c.JSON(error_response.Code, error_response)
}
