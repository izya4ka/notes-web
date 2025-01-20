package database

import (
	"context"
	"errors"

	"github.com/izya4ka/notes-web/user-service/myerrors"
	"github.com/redis/go-redis/v9"
)

func ValidateToken(rdb *redis.Client, username string, token string) error {

	ctx := context.Background()
	db_token, err := rdb.Get(ctx, username).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return myerrors.ErrInvalidToken(0)
		}
		return err
	}
	if db_token != token {
		return myerrors.ErrInvalidToken(0)
	}
	return nil

}
