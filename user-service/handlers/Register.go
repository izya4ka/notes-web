package handlers

import (
	"net/http"

	"github.com/izya4ka/notes-web/user-service/database"
	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/usererrors"
	"github.com/izya4ka/notes-web/user-service/util"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Register handles the user registration process.
// It binds the incoming request to a LogPassRequest model and performs validations
// on the provided username and password. If the username is already taken, it returns
// a conflict status. If registration is successful, it generates a token for the new user
// and adds it to the database.
//
// Parameters:
//   - c: The echo.Context object that holds the request and response.
//   - db: The Gorm database connection used for user data operations.
//   - rdb: The Redis client used for caching or session management.
//
// Returns:
//
//	An error if there is a failure in processing the registration, otherwise returns
//	a HTTP response indicating the status of the registration.
func Register(c echo.Context, db *gorm.DB, rdb *redis.Client) error {

	req := new(models.LogPassRequest)
	if err := c.Bind(req); err != nil {
		util.LogDebugf("%s %s error binding: %s", c.Request().Method, c.Path(), err)
		return util.SendErrorResponse(c, err)
	}

	if err := util.CheckRegLogReq(req); err != nil {
		util.LogDebugf("%s %s error checking creds: %s", c.Request().Method, c.Path(), err)
		return util.SendErrorResponse(c, err)
	}

	if err := database.CheckUserExists(c.Request().Context(), db, req.Username); err == nil {
		util.LogDebugf("%s %s error check user exists: %s", c.Request().Method, c.Path(), err)
		return util.SendErrorResponse(c, usererrors.ErrAlreadyExists)
	}

	if err := database.AddUser(c.Request().Context(), db, rdb, req); err != nil {
		util.LogDebugf("%s %s error adding user to database: %s", c.Request().Method, c.Path(), err)
		return util.SendErrorResponse(c, err)
	}

	token, err := database.UpdateToken(c.Request().Context(), db, rdb, req.Username)
	if err != nil {
		util.LogDebugf("%s %s error updating token: %s", c.Request().Method, c.Path(), err)
		return util.SendErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, models.Token{
		Token: token,
	})
}
