package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/usererrors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// DeleteToken removes the specified user's token from the Redis database.
// It takes a redis.Client instance and the username as parameters.
// If the operation fails, it returns an error. The function executes
// the deletion in the context of a background context, ensuring that
// it can run independently of any parent context.
func DeleteToken(base_ctx context.Context, db *gorm.DB, rdb *redis.Client, username string) error {

	user := new(models.UserPostgres)

	ctx, cancel := context.WithTimeout(base_ctx, 5*time.Second)
	defer cancel()

	err := db.WithContext(ctx).Model(user).Select("token").Where("username = ?", username).First(user).Error
	if err != nil {
		log.Println("Error: ", err)
		switch err {
		case context.DeadlineExceeded:
			return usererrors.ErrTimedOut
		case gorm.ErrRecordNotFound:
			return usererrors.ErrUserNotFound
		default:
			return usererrors.ErrInternal
		}
	}

	if _, err := rdb.Del(ctx, user.Token).Result(); err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Println("Error: ", err)
			if errors.Is(err, context.DeadlineExceeded) {
				return usererrors.ErrTimedOut
			}
			return usererrors.ErrInternal
		}
	}
	return nil
}
