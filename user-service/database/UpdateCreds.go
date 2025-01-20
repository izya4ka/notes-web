package database

import (
	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/util"
	"gorm.io/gorm"
)

// UpdateCreds updates the username and password of a user in the database.
// It takes a database connection and a request containing the new credentials.
// If the provided password cannot be hashed, an error is returned.
// The function attempts to update the user's credentials while ensuring that only
// the specified fields are modified and returns any errors encountered during the process.
func UpdateCreds(db *gorm.DB, req *models.UserChangeCredsRequest) error {

	user := new(models.UserPostgres)
	user.Username = req.NewUsername

	// Hash the new password for security purposes
	new_password, err := util.Hash(req.NewPassword)

	if err != nil {
		return err
	}

	user.Password = new_password

	// Update the user's credentials in the database
	if err := db.Model(user).Select("username", "password").Where("username = ?", req.Username).Updates(user).Error; err != nil {
		return err
	}

	return nil
}
