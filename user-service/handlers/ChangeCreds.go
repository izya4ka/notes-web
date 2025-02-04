package handlers

import (
	"errors"
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
		return util.SendErrorResponse(c, http.StatusBadRequest, "BAD_REQUEST", "Bad data in request!")
	}

	if err := util.CheckRegLogReq(req); err != nil {
		return util.SendErrorResponse(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
	}

	token, err := util.UnrawToken(c.Request().Header.Get("Authorization"))
	if err != nil {
		return util.SendErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", err.Error())
	}

	username, uerr := database.GetUsernameByToken(c, rdb, token)
	if uerr != nil {
		if errors.Is(uerr, usererrors.ErrTimedOut) {
			return util.SendErrorResponse(c, http.StatusRequestTimeout, "REQUEST_TIMEOUT", uerr.Error())
		}
		return util.SendErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", uerr.Error())
	}

	if err := database.CheckUserExists(c, db, req.Username); err == nil && username != req.Username {
		return util.SendErrorResponse(c, http.StatusConflict, "CONFLICT", "User already exists!")
	} else if err != nil {
		if errors.Is(uerr, usererrors.ErrTimedOut) {
			return util.SendErrorResponse(c, http.StatusRequestTimeout, "REQUEST_TIMEOUT", err.Error())
		}
	}

	if err := database.UpdateCreds(c, db, username, req); err != nil {
		if errors.Is(uerr, usererrors.ErrTimedOut) {
			return util.SendErrorResponse(c, http.StatusRequestTimeout, "REQUEST_TIMEOUT", err.Error())
		}
		return util.SendErrorResponse(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
	}
	new_token, terr := database.UpdateToken(c, db, rdb, req.Username)
	if terr != nil {
		if errors.Is(uerr, usererrors.ErrTimedOut) {
			return util.SendErrorResponse(c, http.StatusRequestTimeout, "REQUEST_TIMEOUT", terr.Error())
		}
		return util.SendErrorResponse(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", terr.Error())
	}
	return c.JSON(http.StatusOK, models.Token{
		Token: new_token,
	})
}
