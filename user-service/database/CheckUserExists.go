package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/usererrors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CheckUserExists checks if a user with the given username exists in the database.
// It queries the database for a user matching the provided username and returns an error
// if the user is not found. In case of a successful query, it returns nil indicating
// that the user exists.
//
// Parameters:
// - db: The GORM database instance to perform the query.
// - username: The username of the user to be checked.
//
// Returns:
// An error indicating whether the user was found or not. If the user does not exist,
// it returns a user-specific error indicating that the user was not found. Other errors returned
func CheckUserExists(c echo.Context, db *gorm.DB, username string) error {

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	user := new(models.UserPostgres)
	if err := db.WithContext(ctx).Model(user).Select("username").Where("username = ?", username).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return usererrors.ErrUserNotFound(username)
		}
		log.Println("Error: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			return usererrors.ErrTimedOut
		}
		return usererrors.ErrInternal
	}
	return nil
}
