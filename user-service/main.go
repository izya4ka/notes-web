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

// main initializes the Echo server, connects to the PostgreSQL database,
// and sets up the Redis client. It also registers the necessary HTTP handlers
// for user registration, login, and credential changes. If there is an error
// in connecting to the database, it will print the error and terminate the application.
func main() {
	e := echo.New()

	// Establish a connection to the PostgreSQL database using the provided DB_URL environment variable.
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println("DB success!")

	// Automatically migrate the UserPostgres and Note models to the database.
	db.AutoMigrate(&models.UserPostgres{}, &models.Note{})

	// Create a new Redis client, connecting to the specified Redis server.
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:" + os.Getenv("REDIS_PORT"),
		Password: "", // Password should be set if Redis requires authentication.
		DB:       0,  // Use default DB.
	})

	// Register the POST handler for user registration.
	e.POST("/register", func(c echo.Context) error {
		return handlers.Register(c, db, rdb)
	})

	// Register the POST handler for user login.
	e.POST("/login", func(c echo.Context) error {
		return handlers.Login(c, db, rdb)
	})

	// Register the PUT handler for changing user credentials.
	e.PUT("/change", func(c echo.Context) error {
		return handlers.ChangeCreds(c, db, rdb)
	})

	// Start the Echo server on port 8080, logging fatal errors if they occur.
	e.Logger.Fatal(e.Start(":8080"))
}
