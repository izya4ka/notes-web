package util

import (
	"time"

	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/labstack/echo/v4"
)

func SendErrorResponse(c echo.Context, status_code int, error_type string, message string) error {
	return c.JSON(status_code, models.Error{
		Code:      status_code,
		Error:     error_type,
		Message:   message,
		Path:      c.Path(),
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	})
}
