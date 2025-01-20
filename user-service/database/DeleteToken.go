package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func DeleteToken(rdb *redis.Client, username string) error {
	ctx := context.Background()
	_, err := rdb.Del(ctx, username).Result()
	return err
}
