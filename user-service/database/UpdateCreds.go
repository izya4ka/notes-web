package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/usererrors"
	"github.com/izya4ka/notes-web/user-service/util"
	"gorm.io/gorm"
)

// UpdateCreds updates the username and password of a user in the database.
// It takes a database connection and a request containing the new credentials.
// If the provided password cannot be hashed, an error is returned.
// The function attempts to update the user's credentials while ensuring that only
// the specified fields are modified and returns any errors encountered during the process.
func UpdateCreds(base_ctx context.Context, db *gorm.DB, username string, req *models.LogPassRequest) error {

	user := new(models.UserPostgres)
	user.Username = req.Username

	// Hash the new password for security purposes
	new_password, err := util.Hash(req.Password)

	if err != nil {
		log.Println("Error: ", err)
		return usererrors.ErrInternal
	}

	user.Password = new_password

	ctx, cancel := context.WithTimeout(base_ctx, 5*time.Second)
	defer cancel()

	// Update the user's credentials in the database
	if err := db.WithContext(ctx).Model(user).Select("username", "password").Where("username = ?", username).Updates(user).Error; err != nil {
		log.Println("Error: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			return usererrors.ErrTimedOut
		}
		return usererrors.ErrInternal
	}

	return nil
}
