package database

import (
	"strings"

	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/myerrors"
	"github.com/izya4ka/notes-web/user-service/util"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func AddUser(db *gorm.DB, rdb *redis.Client, req *models.LogPassReq) (string, error) {

	if strings.ContainsAny(req.Username, "\"/!@#$%^&*()+=[]{}';:?*") {
		return "", myerrors.ErrNotWithoutSpecSym(req.Username)
	}

	password, perr := util.Hash(req.Password)

	if perr != nil {
		return "", perr
	}

	user := models.UserPostgres{
		Username: req.Username,
		Password: password,
	}

	result := db.Create(&user)

	token, terr := UpdateToken(db, rdb, req.Username)
	if terr != nil {
		return "", terr
	}

	if result.Error != nil {
		return "", result.Error
	}

	return token, nil
}
