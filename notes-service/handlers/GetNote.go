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

func GetNote(c *gin.Context, db *gorm.DB, username string) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		util.SendErrorResponse(c, http.StatusBadRequest, "BAD_REQUEST", "bad request")
		return
	}

	note, derr := database.GetNote(c.Request.Context(), db, username, id)
	if derr != nil {
		switch derr {
		case noteserrors.ErrTimedOut:
			util.SendErrorResponse(c, http.StatusRequestTimeout, "REQUEST_TIMEOUT", derr.Error())
		case noteserrors.ErrNotFound:
			util.SendErrorResponse(c, http.StatusNotFound, "NOT_FOUND", derr.Error())
		default:
			util.SendErrorResponse(c, http.StatusInternalServerError, "INTERNAL_ERROR", derr.Error())
		}
		return
	}

	c.JSON(http.StatusOK, note)
}
