package database

import (
	"context"
	"errors"

	"github.com/izya4ka/notes-web/user-service/usererrors"
	"github.com/redis/go-redis/v9"
)

func GetUsernameByToken(rdb *redis.Client, token string) (string, error) {
	ctx := context.Background()
	username, err := rdb.Get(ctx, token).Result()
	if errors.Is(err, redis.Nil) {
		return "", usererrors.ErrInvalidToken(0)
	}
	return username, nil
}
