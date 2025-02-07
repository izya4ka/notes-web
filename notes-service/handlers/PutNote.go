package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/database"
	"github.com/izya4ka/notes-web/notes-service/models"
	"github.com/izya4ka/notes-web/notes-service/util"
	"gorm.io/gorm"
)

func PutNode(c *gin.Context, db *gorm.DB, username string) {
	var req models.BaseNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendErrorResponse(c, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		util.SendErrorResponse(c, err)
		return
	}

	note, nerr := database.PutNote(c.Request.Context(), db, &req, username, id)
	if nerr != nil {
		util.SendErrorResponse(c, nerr)
		return
	}

	c.JSON(http.StatusOK, note)
}
