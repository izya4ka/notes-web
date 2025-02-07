package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/database"
	"github.com/izya4ka/notes-web/notes-service/util"
	"gorm.io/gorm"
)

func GetNote(c *gin.Context, db *gorm.DB, username string) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		util.SendErrorResponse(c, err)
		return
	}

	note, derr := database.GetNote(c.Request.Context(), db, username, id)
	if derr != nil {
		util.SendErrorResponse(c, derr)
		return
	}

	c.JSON(http.StatusOK, note)
}
