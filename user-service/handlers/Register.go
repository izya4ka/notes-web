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

	if err := database.CheckUserExists(db, req.Username); err == nil {
		return c.JSON(http.StatusConflict, models.Error{
			Code:      http.StatusConflict,
			Error:     "CONFLICT",
			Message:   "User already exists!",
			Timestamp: current_time,
			Path:      c.Path(),
		})
	}

	if err := database.AddUser(db, rdb, req); err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Code:      http.StatusInternalServerError,
			Error:     "INTERNAL_ERROR",
			Message:   "Error has occured on the server side",
			Timestamp: current_time,
			Path:      c.Path(),
		})
	}

	token, err := database.UpdateToken(db, rdb, req.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Code:      http.StatusInternalServerError,
			Error:     "INTERNAL_ERROR",
			Message:   "Error has occured on the server side",
			Timestamp: current_time,
			Path:      c.Path(),
		})
	}

	return c.String(http.StatusCreated, token)
}
