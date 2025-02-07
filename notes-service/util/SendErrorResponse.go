package util

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/models"
	"github.com/izya4ka/notes-web/notes-service/noteserrors"
)

func SendErrorResponse(c *gin.Context, err error) {

	var error_response models.Error

	switch err {
	case noteserrors.ErrInvalidToken:
		error_response.Code = http.StatusUnauthorized
	case noteserrors.ErrInternal:
		error_response.Code = http.StatusInternalServerError
	case noteserrors.ErrTimedOut:
		error_response.Code = http.StatusRequestTimeout
	case noteserrors.ErrNotFound:
		error_response.Code = http.StatusNotFound
	}

	error_response.Error = http.StatusText(error_response.Code)
	error_response.Message = err.Error()
	error_response.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	error_response.Path = c.Request.URL.Path

	c.JSON(error_response.Code, error_response)
}
