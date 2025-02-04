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
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err := util.CheckRegLogReq(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	token, err := util.UnrawToken(c.Request().Header.Get("Authorization"))
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	username, uerr := database.GetUsernameByToken(rdb, token)
	if uerr != nil {
		return c.String(http.StatusUnauthorized, uerr.Error())
	}

	if err := database.CheckUserExists(db, req.Username); err == nil && username != req.Username {
		return c.String(http.StatusConflict, usererrors.ErrAlreadyExists(req.Username).Error())
	}

	if err := database.UpdateCreds(db, username, req); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	new_token, terr := database.UpdateToken(db, rdb, req.Username)
	if terr != nil {
		return c.String(http.StatusInternalServerError, terr.Error())
	}

	return c.String(http.StatusOK, new_token)
}
