package handlers

import (
	"net/http"

	"github.com/izya4ka/notes-web/user-service/database"
	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/util"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// PostLogin handles user login requests. It processes the login credentials
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
func PostLogin(c echo.Context, db *gorm.DB, rdb *redis.Client) error {

	req := new(models.LogPassRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err := util.CheckRegLogReq(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := database.CheckPassword(req, db); err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	token, terr := database.UpdateToken(rdb, req.Username)
	if terr != nil {
		return c.String(http.StatusInternalServerError, terr.Error())
	}

	return c.String(http.StatusOK, token)
}
