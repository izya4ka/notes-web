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

func PutUpdate(c echo.Context, db *gorm.DB, rdb *redis.Client) error {
	req := new(models.LogPassReq)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	token, err := util.UnrawToken(c.Request().Header.Get("Authorization"))
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}
	username, terr := database.GetUsernameByToken(rdb, token)
	if terr != nil {
		return c.String(http.StatusUnauthorized, terr.Error())
	}	
	
	if err := database.CheckUserExists(db, username); err != nil {
		return err
	}

	if err := database.UpdateUser(db, username, req); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	new_token, terr := database.SetToken(db, rdb, username)
	if terr != nil {
		return c.String(http.StatusInternalServerError, terr.Error())
	}

	return c.String(http.StatusOK, new_token)
}
