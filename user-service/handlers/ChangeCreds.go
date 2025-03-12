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

// ChangeCreds handles the request to change user credentials.
// It processes the incoming request to update the user's username and password.
// The function performs validation checks to ensure the request is valid,
// including checking if the username already exists and validating the user's token.
// In case of any conflicts or errors during the process, appropriate HTTP status codes
// are returned along with error messages.
// If the change is successful, a new token is generated and returned to the user.
//
// Parameters:
// - c: The Echo context which carries request and response information.
// - db: The Gorm database instance used for querying the database.
// - rdb: The Redis client used for token validation and management.
//
// Returns:
// - An error if the process encounters any issues, otherwise returns nil.
func ChangeCreds(c echo.Context, db *gorm.DB, rdb *redis.Client) error {

	req := new(models.LogPassRequest)
	if err := c.Bind(req); err != nil {
		util.LogDebugf("%s %s error binding: %s", c.Request().Method, c.Path(), err)
		return util.SendErrorResponse(c, err)
	}

	if err := util.CheckRegLogReq(req); err != nil {
		util.LogDebugf("%s %s checking creds: %s", c.Request().Method, c.Path(), err)
		return util.SendErrorResponse(c, err)
	}

	token, err := util.UnrawToken(c.Request().Header.Get("Authorization"))
	if err != nil {
		util.LogDebugf("%s %s error unraw token: %s", c.Request().Method, c.Path(), err)
		return util.SendErrorResponse(c, err)
	}

	username, uerr := database.GetUsername(c.Request().Context(), rdb, token)
	if uerr != nil {
		util.LogDebugf("%s %s error getting username: %s", c.Request().Method, c.Path(), uerr)
		return util.SendErrorResponse(c, uerr)
	}

	if err := database.CheckUserExists(c.Request().Context(), db, req.Username); err == nil && username != req.Username {
		util.LogDebugf("%s %s error check user exists: %s", c.Request().Method, c.Path(), err)
		return util.SendErrorResponse(c, usererrors.ErrAlreadyExists)
	} else if err != nil {
		util.LogDebugf("%s %s user not exists: %s", c.Request().Method, c.Path(), err)
		return util.SendErrorResponse(c, err)
	}

	if err := database.UpdateCreds(c.Request().Context(), db, username, req); err != nil {
		util.LogDebugf("%s %s error updating credentials: %s", c.Request().Method, c.Path(), err)
		return util.SendErrorResponse(c, err)
	}

	new_token, terr := database.UpdateToken(c.Request().Context(), db, rdb, req.Username)
	if terr != nil {
		util.LogDebugf("%s %s error updating token: %s", c.Request().Method, c.Path(), terr)
		return util.SendErrorResponse(c, terr)
	}

	return c.JSON(http.StatusOK, models.Token{
		Token: new_token,
	})
}
