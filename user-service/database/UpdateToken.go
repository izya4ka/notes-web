package database

import (
	"context"
	"time"

	"github.com/izya4ka/notes-web/user-service/util"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func UpdateToken(db *gorm.DB, rdb *redis.Client, username string) (string, error) {

	token, jerr := util.CalcToken(username)
	if jerr != nil {
		return "", jerr
	}

	ctx := context.Background()
	rdb.Del(ctx, username)
	go rdb.Set(ctx, username, token, time.Hour*24*7)
	return token, nil
}
