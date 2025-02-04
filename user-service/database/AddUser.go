package database

import (
	"log"
	"strings"

	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/usererrors"
	"github.com/izya4ka/notes-web/user-service/util"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// AddUser adds a new user to the PostgreSQL database and generates a token for the user.
// It validates the username to ensure it does not contain special characters,
// hashes the provided password, and creates a user record in the database.
// If successful, it returns the generated token; otherwise, it returns an error.
//
// Parameters:
// - db: Database connection instance to perform operations on the PostgreSQL database.
// - rdb: Redis client instance to manage user sessions.
// - req: A pointer to a LogPassRequest containing the username and password for the new user.
//
// Returns:
// - A string containing the generated token if successful, or an empty string if an error occurs.
// - An error indicating what went wrong during the user creation process.
func AddUser(db *gorm.DB, rdb *redis.Client, req *models.LogPassRequest) error {

	// Check if the username contains any special characters
	if strings.ContainsAny(req.Username, "\"/!@#$%^&*()+=[]{}';:?*") {
		return usererrors.ErrNotWithoutSpecSym(req.Username)
	}

	// Hash the user's password
	password, perr := util.Hash(req.Password)

	if perr != nil {
		log.Println("Error: ", perr)
		return usererrors.ErrInternal
	}

	// Create a new user model instance with the provided username and hashed password
	user := models.UserPostgres{
		Username: req.Username,
		Password: password,
	}

	// Insert the new user into the database
	result := db.Create(&user)

	// Check if there was an error during the database operation
	if result.Error != nil {
		log.Println("Error: ", result.Error)
		return usererrors.ErrInternal
	}

	// Return the generated token
	return nil
}
