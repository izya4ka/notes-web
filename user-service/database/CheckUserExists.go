package database

import (
	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/myerrors"
	"gorm.io/gorm"
)

func CheckUserExists(db *gorm.DB, username string) error {
	user := new(models.UserPostgres)
	if err := db.Model(user).Select("username").Where("username = ?", username).First(user).Error; err != nil {
		return myerrors.ErrUserNotFound(username)
	}
	return nil
}