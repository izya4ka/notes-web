package database

import (
	"context"
	"errors"

	"github.com/izya4ka/notes-web/user-service/usererrors"
	"github.com/redis/go-redis/v9"
)

// ValidateToken checks if the provided token matches the token stored in Redis for the given username.
// It returns an error if the username does not exist or if the tokens do not match.
// If the username is not found in the database, it returns a user-specific error indicating an invalid token.
// It uses a context to interact with the Redis client and handles potential errors from the Redis operations.
func ValidateToken(rdb *redis.Client, username string, token string) error {
	ctx := context.Background()
	db_token, err := rdb.Get(ctx, username).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return usererrors.ErrInvalidToken(0) // Token not found for the specified username.
		}
		return err // Return any other Redis error encountered.
	}
	if db_token != token {
		return usererrors.ErrInvalidToken(0) // Tokens do not match; return invalid token error.
	}
	return nil // Valid token.
}
