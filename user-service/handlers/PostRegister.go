package handlers

import (
	"net/http"

	"github.com/izya4ka/notes-web/user-service/database"
	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/myerrors"
	"github.com/izya4ka/notes-web/user-service/util"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func PostRegister(c echo.Context, db *gorm.DB, rdb *redis.Client) error {
	
	req := new(models.LogPassReq)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err := util.CheckRegLogReq(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := database.CheckUserExists(db, req.Username); err == nil {
		return c.String(http.StatusConflict, myerrors.ErrAlreadyExists(req.Username).Error())
	}

	token, err := database.AddUser(db, rdb, req)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusCreated, token)
}
