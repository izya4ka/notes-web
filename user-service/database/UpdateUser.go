package database

import (
	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/util"
	"gorm.io/gorm"
)

func UpdateUser(db *gorm.DB, username string, req *models.LogPassReq) error {

	user := new(models.UserPostgres)
	user.Username = req.Username
	new_password, err := util.Hash(req.Password)
	if err != nil {
		return err
	}
	user.Password = new_password

	if err := db.Model(user).Select("username", "password").Where("username = ?", username).Updates(user).Error; err != nil {
		return err
	}

	return nil
}