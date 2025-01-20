package database

import (
	"context"
	"time"

	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/util"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetToken(db *gorm.DB, rdb *redis.Client, username string) (string, error) {
	token, jerr := util.CalcToken(username)

	if jerr != nil {
		return "", jerr
	}

	var old_token string

	if err := db.Model(&models.UserPostgres{}).Select("username", "current_token").Where("username = ?", username).Pluck("current_token", &old_token).Error; err != nil {
		return "", err
	}

	if err := db.Model(&models.UserPostgres{}).Select("username", "current_token").Where("username = ?", username).Update("current_token", token).Error; err != nil {
		return "", err
	}

	ctx := context.Background()
	go rdb.Del(ctx, old_token)
	go rdb.Set(ctx, token, username, time.Hour * 24 * 7)
	return token, nil
}
