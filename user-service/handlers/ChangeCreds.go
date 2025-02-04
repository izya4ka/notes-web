package handlers

import (
	"net/http"
	"time"

	"github.com/izya4ka/notes-web/user-service/database"
	"github.com/izya4ka/notes-web/user-service/models"
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

	current_time := time.Now().Format("2006-01-02 15:04:05")

	req := new(models.LogPassRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{
			Code:      http.StatusBadRequest,
			Error:     "BAD_REQUEST",
			Message:   "bad data in request",
			Timestamp: current_time,
			Path:      c.Path(),
		})
	}

	if err := util.CheckRegLogReq(req); err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{
			Code:      http.StatusBadRequest,
			Error:     "BAD_REQUEST",
			Message:   err.Error(),
			Timestamp: current_time,
			Path:      c.Path(),
		})
	}

	token, err := util.UnrawToken(c.Request().Header.Get("Authorization"))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Error{
			Code:      http.StatusUnauthorized,
			Error:     "UNAUTHORIZED",
			Message:   err.Error(),
			Timestamp: current_time,
			Path:      c.Path(),
		})
	}

	username, uerr := database.GetUsernameByToken(rdb, token)
	if uerr != nil {
		return c.JSON(http.StatusUnauthorized, models.Error{
			Code:      http.StatusUnauthorized,
			Error:     "UNAUTHORIZED",
			Message:   uerr.Error(),
			Timestamp: current_time,
			Path:      c.Path(),
		})
	}

	if err := database.CheckUserExists(db, req.Username); err == nil && username != req.Username {
		return c.JSON(http.StatusConflict, models.Error{
			Code:      http.StatusConflict,
			Error:     "CONFLICT",
			Message:   "User already exists!",
			Timestamp: current_time,
			Path:      c.Path(),
		})
	}

	if err := database.UpdateCreds(db, username, req); err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Code:      http.StatusInternalServerError,
			Error:     "INTERNAL_ERROR",
			Message:   "Error has occured on the server side",
			Timestamp: current_time,
			Path:      c.Path(),
		})
	}
	new_token, terr := database.UpdateToken(db, rdb, req.Username)
	if terr != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Code:      http.StatusInternalServerError,
			Error:     "INTERNAL_ERROR",
			Message:   "Error has occured on the server side",
			Timestamp: current_time,
			Path:      c.Path(),
		})
	}
	return c.String(http.StatusOK, new_token)
}
