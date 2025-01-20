package database

import (
	"context"
	"time"

	"github.com/izya4ka/notes-web/user-service/util"
	"github.com/redis/go-redis/v9"
)

// UpdateToken generates a new token for the given username and updates it in
// the Redis. It first calculates the token using the utility function
// CalcToken. If the token calculation fails, it returns an error.
// It then attempts to delete any existing token associated with the username in Redis
// before setting the newly generated token with a TTL of one week.
// If any Redis operations fail, it returns the error encountered.
//
// This function requires a Redis client (redis.Client).
//
// Parameters:
// - rdb: the Redis client used to store and manage tokens.
// - username: the username for which the token is being updated.
//
// Returns:
// - A string containing the new token if the operation succeeds, or an error if it fails.
func UpdateToken(rdb *redis.Client, username string) (string, error) {

	token, jerr := util.CalcToken(username)
	if jerr != nil {
		return "", jerr
	}

	ctx := context.Background()
	if err := DeleteToken(rdb, username); err != nil {
		return "", err
	}
	if _, err := rdb.Set(ctx, username, token, time.Hour*24*7).Result(); err != nil {
		return "", err
	}
	return token, nil
}
