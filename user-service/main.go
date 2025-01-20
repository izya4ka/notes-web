package main

import (
	"fmt"
	"os"

	"github.com/izya4ka/notes-web/user-service/handlers"
	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println("DB success!")

	db.AutoMigrate(&models.UserPostgres{}, &models.Note{})

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	e.POST("/register", func(c echo.Context) error {
		return handlers.PostRegister(c, db, rdb)
	})
	e.POST("/login", func(c echo.Context) error {
		return handlers.PostLogin(c, db, rdb)
	})
	e.PUT("/change", func(c echo.Context) error {
		return handlers.PutUpdate(c, db, rdb)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
