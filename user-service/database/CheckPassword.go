package database

import (
	"errors"

	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/myerrors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckPassword(req *models.LogPassReq, db *gorm.DB) error {

	user := new(models.UserPostgres)

	err := db.Model(user).Where("username = ?", req.Username).Select("username", "password").First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerrors.ErrUserNotFound(req.Username)
		} else {
			return err
		}
	}

	herr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if herr != nil {
		if errors.Is(herr, bcrypt.ErrMismatchedHashAndPassword) {
			return myerrors.ErrMismatchPass(req.Username)
		} else {
			return herr
		}
	}
	return nil
}
