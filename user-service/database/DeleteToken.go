package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// DeleteToken removes the specified user's token from the Redis database.
// It takes a redis.Client instance and the username as parameters.
// If the operation fails, it returns an error. The function executes
// the deletion in the context of a background context, ensuring that
// it can run independently of any parent context.
func DeleteToken(rdb *redis.Client, username string) error {
	ctx := context.Background()
	_, err := rdb.Del(ctx, username).Result()
	return err
}
