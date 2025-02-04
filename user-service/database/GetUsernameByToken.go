package database

import (
	"context"
	"errors"
	"log"

	"github.com/izya4ka/notes-web/user-service/usererrors"
	"github.com/redis/go-redis/v9"
)

func GetUsernameByToken(rdb *redis.Client, token string) (string, error) {
	ctx := context.Background()
	username, err := rdb.Get(ctx, token).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", usererrors.ErrInvalidToken
		}
		log.Println("Error: ", err)
		return "", usererrors.ErrInternal
	}
	return username, nil
}
