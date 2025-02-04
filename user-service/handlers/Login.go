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

// Login handles user login requests. It processes the login credentials
// provided by the user, validates them, and returns an authentication token on success.
//
// The function expects a JSON request body containing the login information,
// which is bound to the LogPassRequest model. If the binding fails or the
// request validation fails, it responds with a Bad Request status.
//
// After successful validation, it checks the user's password against the database.
// If the password is incorrect, it responds with an Unauthorized status.
//
// On successful login, it updates the user's authentication token in the
// Redis, returning the token with an OK status. If there are issues updating
// the token, it responds with an Internal Server Error status.
//
// Parameters:
// - c: The echo.Context containing the HTTP request and response.
// - db: The gorm.DB instance for database operations.
// - rdb: The redis.Client instance for Redis operations.
//
// Returns:
// - error: An error, if any occurred during the login process. Otherwise, it returns nil.
func Login(c echo.Context, db *gorm.DB, rdb *redis.Client) error {

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

	if err := database.CheckPassword(req, db); err != nil {
		return c.JSON(http.StatusUnauthorized, models.Error{
			Code:      http.StatusUnauthorized,
			Error:     "UNAUTHORIZED",
			Message:   err.Error(),
			Timestamp: current_time,
			Path:      c.Path(),
		})
	}

	token, terr := database.UpdateToken(db, rdb, req.Username)
	if terr != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Code:      http.StatusInternalServerError,
			Error:     "INTERNAL_ERROR",
			Message:   "Error has occured on the server side",
			Timestamp: current_time,
			Path:      c.Path(),
		})
	}

	return c.String(http.StatusOK, token)
}
