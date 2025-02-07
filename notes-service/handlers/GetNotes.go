package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/database"
	"github.com/izya4ka/notes-web/notes-service/models"
	"github.com/izya4ka/notes-web/notes-service/noteserrors"
	"github.com/izya4ka/notes-web/notes-service/util"
	"gorm.io/gorm"
)

func GetNotes(c *gin.Context, db *gorm.DB, username string) {
	var req models.GetNotesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendErrorResponse(c, http.StatusBadRequest, "BAD_REQUEST", "bad request")
		return
	}

	notes, err := database.GetNotes(c.Request.Context(), db, username, req.Amount)
	if err != nil {
		if errors.Is(err, noteserrors.ErrTimedOut) {
			util.SendErrorResponse(c, http.StatusRequestTimeout, "REQUEST_TIMEOUT", err.Error())
			return
		}
		util.SendErrorResponse(c, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, notes)
}
