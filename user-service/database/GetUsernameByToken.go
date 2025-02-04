package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/izya4ka/notes-web/user-service/usererrors"
	"github.com/redis/go-redis/v9"
)

func GetUsernameByToken(rdb *redis.Client, token string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	username, err := rdb.Get(ctx, token).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", usererrors.ErrInvalidToken
		}
		log.Println("Error: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			return "", usererrors.ErrTimedOut
		}
		return "", usererrors.ErrInternal
	}
	return username, nil
}
