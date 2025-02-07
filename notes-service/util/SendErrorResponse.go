package util

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/models"
)

func SendErrorResponse(c *gin.Context, status_code int, error_type string, message string) {
	c.JSON(status_code, models.Error{
		Code:      status_code,
		Error:     error_type,
		Message:   message,
		Path:      c.Request.URL.Path,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	})
}
