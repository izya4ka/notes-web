package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/database"
	"github.com/izya4ka/notes-web/notes-service/util"
	"gorm.io/gorm"
)

func GetNotes(c *gin.Context, db *gorm.DB) {
	username := c.GetHeader("Username")

	amountStr := c.DefaultQuery("amount", "5")
	pageStr := c.DefaultQuery("page", "0")

	util.LogDebugf("%s %s", c.Request.Method, c.Request.RequestURI)
	util.LogDebugf("amountStr: %s", amountStr)
	util.LogDebugf("pageStr: %s", pageStr)

	amount, aerr := strconv.Atoi(amountStr)
	if aerr != nil {
		util.LogDebugf("%s %s query: %s", c.Request.Method, c.Request.URL.Path, aerr)
		util.SendErrorResponse(c, aerr)
		return
	}

	page, perr := strconv.Atoi(pageStr)
	if perr != nil {
		util.LogDebugf("%s %s query: %s", c.Request.Method, c.Request.URL.Path, perr)
		util.SendErrorResponse(c, perr)
		return
	}

	notes, err := database.GetNotes(c.Request.Context(), db, username, amount, page)
	if err != nil {
		util.LogDebugf("%s %s database get notes: %s", c.Request.Method, c.Request.URL.Path, perr)
		util.SendErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, notes)
}
