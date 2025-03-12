package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/database"
	"github.com/izya4ka/notes-web/notes-service/util"
	"gorm.io/gorm"
)

func GetNote(c *gin.Context, db *gorm.DB) {
	username := c.GetHeader("Username")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		util.LogDebugf("%s %s param: %s", c.Request.Method, c.Request.URL.Path, err)
		util.SendErrorResponse(c, err)
		return
	}

	note, derr := database.GetNote(c.Request.Context(), db, username, id)
	if derr != nil {
		util.LogDebugf("%s %s database get note: %s", c.Request.Method, c.Request.URL.Path, err)
		util.SendErrorResponse(c, derr)
		return
	}

	c.JSON(http.StatusOK, note)
}
