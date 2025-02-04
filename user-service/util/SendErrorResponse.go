package util

import (
	"net/http"
	"time"

	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/labstack/echo/v4"
)

func SendErrorResponse(c echo.Context, status_code int, error_type, message string) error {
	var error_struct models.Error
	error_struct.Path = c.Path()
	error_struct.Timestamp = time.Now().Format("2006-01-02 15:04:05")

	if status_code == http.StatusInternalServerError {
		error_struct.Error = "INTERNAL_SERVER_ERROR"
		error_struct.Code = http.StatusInternalServerError
		error_struct.Message = "Error occured on the server side!"
		return c.JSON(status_code, error_struct)
	}

	error_struct.Error = error_type
	error_struct.Code = status_code
	error_struct.Message = message
	return c.JSON(status_code, error_struct)
}
