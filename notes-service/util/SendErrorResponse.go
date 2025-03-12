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

	error_response.Message = err.Error()

	switch err {
	case noteserrors.ErrInvalidToken:
		error_response.Code = http.StatusUnauthorized
	case noteserrors.ErrInternal:
		error_response.Code = http.StatusInternalServerError
	case noteserrors.ErrTimedOut:
		error_response.Code = http.StatusRequestTimeout
	case noteserrors.ErrNotFound:
		error_response.Code = http.StatusNotFound
	case noteserrors.ErrWrongParams:
		error_response.Code = http.StatusBadRequest
	default:
		error_response.Code = http.StatusInternalServerError
		LogErrorf("%s %s %d -> Error: %s", c.Request.Method, c.Request.URL.Path, error_response.Code, err)
		error_response.Message = "Undocumented Error"
	}

	error_response.Error = http.StatusText(error_response.Code)

	error_response.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	error_response.Path = c.Request.URL.Path

	c.JSON(error_response.Code, error_response)
}
