package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/database"
	"github.com/izya4ka/notes-web/notes-service/util"
	"gorm.io/gorm"
)

func DeleteNote(c *gin.Context, db *gorm.DB) {
	username := c.GetHeader("Username")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		util.LogDebugf("%s %s param: %s", c.Request.Method, c.Request.URL.Path, err)
		util.SendErrorResponse(c, err)
		return
	}

	if err := database.DeleteNote(c.Request.Context(), db, username, id); err != nil {
		util.LogDebugf("%s %s database delete note: %s", c.Request.Method, c.Request.URL.Path, err)
		util.SendErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}
