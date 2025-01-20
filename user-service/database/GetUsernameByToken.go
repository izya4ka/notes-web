package database

import (
	"context"
	"errors"

	"github.com/izya4ka/notes-web/user-service/myerrors"
	"github.com/redis/go-redis/v9"
)

func GetUsernameByToken(rdb *redis.Client, token string) (string, error) {

	ctx := context.Background()
	username, err := rdb.Get(ctx, token).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", myerrors.ErrInvalidToken(0)
		} else {
			return "", err
		}
	}
	
	return username, err
}