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
		return util.SendErrorResponse(c, http.StatusBadRequest, "BAD_REQUEST", "Bad data in request!")
	}

	if err := util.CheckRegLogReq(req); err != nil {
		return util.SendErrorResponse(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
	}

	if err := database.CheckUserExists(c, db, req.Username); err == nil {
		return util.SendErrorResponse(c, http.StatusConflict, "CONFLICT", "User "+req.Username+" already exists!")
	} else {
		if errors.Is(err, usererrors.ErrTimedOut) {
			return util.SendErrorResponse(c, http.StatusRequestTimeout, "REQUEST_TIMEOUT", err.Error())
		}
	}

	if err := database.AddUser(c, db, rdb, req); err != nil {
		if errors.Is(err, usererrors.ErrTimedOut) {
			return util.SendErrorResponse(c, http.StatusRequestTimeout, "REQUEST_TIMEOUT", err.Error())
		}
		return util.SendErrorResponse(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
	}

	token, err := database.UpdateToken(c, db, rdb, req.Username)
	if err != nil {
		if errors.Is(err, usererrors.ErrTimedOut) {
			return util.SendErrorResponse(c, http.StatusRequestTimeout, "REQUEST_TIMEOUT", err.Error())
		}
		return util.SendErrorResponse(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
	}

	return c.String(http.StatusCreated, token)
}
