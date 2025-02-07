package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/database"
	"github.com/izya4ka/notes-web/notes-service/noteserrors"
	"github.com/izya4ka/notes-web/notes-service/util"
	"gorm.io/gorm"
)

func DeleteNote(c *gin.Context, db *gorm.DB, username string) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		util.SendErrorResponse(c, http.StatusBadRequest, "BAD_REQUEST", "bad request")
		return
	}

	if err := database.DeleteNote(c.Request.Context(), db, username, id); err != nil {
		switch err {
		case noteserrors.ErrTimedOut:
			util.SendErrorResponse(c, http.StatusRequestTimeout, "REQUEST_TIMEOUT", err.Error())
		case noteserrors.ErrNotFound:
			util.SendErrorResponse(c, http.StatusNotFound, "NOT_FOUND", err.Error())
		default:
			util.SendErrorResponse(c, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		}
		return
	}

	c.Status(http.StatusOK)
}
