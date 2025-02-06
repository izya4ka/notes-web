package main

import (
	"context"
	"net"
	"os"

	"log"

	pb "github.com/izya4ka/notes-web/user-service/proto"

	"github.com/izya4ka/notes-web/user-service/handlers"
	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/userrpc"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
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
		log.Fatal(err)
	}
	log.Print("DB connection success!")

	// Automatically migrate the UserPostgres and Note models to the database.
	db.AutoMigrate(&models.UserPostgres{}, &models.Note{})

	// Create a new Redis client, connecting to the specified Redis server.
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:" + os.Getenv("REDIS_PORT"),
		Password: "", // Password should be set if Redis requires authentication.
		DB:       0,  // Use default DB.
	})

	// Check if Redis connection established
	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatalf("Failed to listen for RPC: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTokenServiceServer(s, userrpc.NewRPCServer(rdb))

	log.Printf("RPC server listening at %v", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("RPC server failed to serve: %v", err)
		}
	}()

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
	e.Logger.Fatal(e.Start(":" + os.Getenv("USER_SERVICE_PORT")))
}
