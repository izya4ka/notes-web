package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/database"
	"github.com/izya4ka/notes-web/notes-service/models"
	"github.com/izya4ka/notes-web/notes-service/util"
	"gorm.io/gorm"
)

func GetNotes(c *gin.Context, db *gorm.DB, username string) {
	var req models.GetNotesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendErrorResponse(c, err)
		return
	}

	notes, err := database.GetNotes(c.Request.Context(), db, username, req.Amount)
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, notes)
}
