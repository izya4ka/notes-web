package handlers

import (
	"net/http"

	"github.com/izya4ka/notes-web/user-service/database"
	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/util"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func PostLogin(c echo.Context, db *gorm.DB, rdb *redis.Client) error {

	req := new(models.LogPassReq)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err := util.CheckRegLogReq(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := database.CheckPassword(req, db); err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	token, terr := database.UpdateToken(db, rdb, req.Username)
	if terr != nil {
		return c.String(http.StatusInternalServerError, terr.Error())
	}

	return c.String(http.StatusOK, token)
}
