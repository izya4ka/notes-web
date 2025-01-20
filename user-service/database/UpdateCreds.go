package database

import (
	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/util"
	"gorm.io/gorm"
)

func UpdateCreds(db *gorm.DB, req *models.UserChangeCreds) error {

	user := new(models.UserPostgres)
	user.Username = req.NewUsername
	new_password, err := util.Hash(req.NewPassword)

	if err != nil {
		return err
	}

	user.Password = new_password

	if err := db.Model(user).Select("username", "password").Where("username = ?", req.Username).Updates(user).Error; err != nil {
		return err
	}

	return nil
}
