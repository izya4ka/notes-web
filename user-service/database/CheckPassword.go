package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/usererrors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CheckPassword verifies that the provided password matches the stored password
// for a user identified by the given username in the database. It takes a
// LogPassRequest containing the username and password, and a database
// connection. If the user is not found, it returns an error indicating the
// user does not exist. If the password does not match, it returns an error
// indicating a password mismatch. Any other errors encountered during
// the database operation or password comparison are also returned.
//
// Parameters:
//   - req: A pointer to a LogPassRequest struct containing the username
//     and password to check.
//   - db: A pointer to a gorm.DB instance for database operations.
//
// Returns:
//   - An error if the user is not found, the password does not match, or
//     if any other error occurs. Returns nil if the password is verified
//     successfully.
func CheckPassword(c echo.Context, req *models.LogPassRequest, db *gorm.DB) error {
	user := new(models.UserPostgres)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := db.WithContext(ctx).Model(user).Where("username = ?", req.Username).Select("username", "password").First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return usererrors.ErrUserNotFound(req.Username)
		} else {
			log.Println("Error: ", err)
			if errors.Is(err, context.DeadlineExceeded) {
				return usererrors.ErrTimedOut
			}
			return usererrors.ErrInternal
		}
	}

	herr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if herr != nil {
		if errors.Is(herr, bcrypt.ErrMismatchedHashAndPassword) {
			return usererrors.ErrMismatchPass(req.Username)
		} else {
			log.Println("Error: ", herr)
			return usererrors.ErrInternal
		}
	}
	return nil
}
