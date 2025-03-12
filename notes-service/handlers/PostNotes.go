package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/database"
	"github.com/izya4ka/notes-web/notes-service/models"
	"github.com/izya4ka/notes-web/notes-service/util"
	"gorm.io/gorm"
)

func PostNotes(c *gin.Context, db *gorm.DB) {
	username := c.GetHeader("Username")
	var req models.BaseNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.LogDebugf("%s %s binding: %s", c.Request.Method, c.Request.URL.Path, err)
		util.SendErrorResponse(c, err)
		return
	}

	note, err := database.AddNote(c.Request.Context(), db, &req, username)
	if err != nil {
		util.LogDebugf("%s %s database add note: %s", c.Request.Method, c.Request.URL.Path, err)
		util.SendErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusAccepted, note)
}
