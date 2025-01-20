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

func PutChangeCreds(c echo.Context, db *gorm.DB, rdb *redis.Client) error {

	req := new(models.UserChangeCreds)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	reglogreq := models.LogPassReq{
		Username: req.NewUsername,
		Password: req.NewPassword,
	}

	if err := util.CheckRegLogReq(&reglogreq); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	token, err := util.UnrawToken(c.Request().Header.Get("Authorization"))
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	if err := database.CheckUserExists(db, req.Username); err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	flag := false
	if req.NewUsername != req.Username {
		flag = true
		if err := database.CheckUserExists(db, req.NewUsername); err == nil {
			return c.String(http.StatusConflict, myerrors.ErrAlreadyExists(req.NewUsername).Error())
		}
	}

	if err := database.ValidateToken(rdb, req.Username, token); err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	if flag {
		if err := database.DeleteToken(rdb, req.Username); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	if err := database.UpdateCreds(db, req); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	new_token, terr := database.UpdateToken(db, rdb, req.NewUsername)
	if terr != nil {
		return c.String(http.StatusInternalServerError, terr.Error())
	}

	return c.String(http.StatusOK, new_token)
}
