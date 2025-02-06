package server

import (
	"context"
	"log"
	"net"

	"github.com/izya4ka/notes-web/user-service/handlers"
	"github.com/izya4ka/notes-web/user-service/models"
	pb "github.com/izya4ka/notes-web/user-service/proto"
	"github.com/izya4ka/notes-web/user-service/userrpc"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	echo        *echo.Echo
	db          *gorm.DB
	rdb         *redis.Client
	grpc_server *grpc.Server
}

func (s *Server) InitEcho(port string) {
	e := echo.New()

	// Register the POST handler for user registration.
	e.POST("/register", func(c echo.Context) error {
		return handlers.Register(c, s.db, s.rdb)
	})

	// Register the POST handler for user login.
	e.POST("/login", func(c echo.Context) error {
		return handlers.Login(c, s.db, s.rdb)
	})

	// Register the PUT handler for changing user credentials.
	e.PUT("/change", func(c echo.Context) error {
		return handlers.ChangeCreds(c, s.db, s.rdb)
	})

	// Start the Echo server on port 8080, logging fatal errors if they occur.
	e.Logger.Fatal(e.Start(":" + port))
	s.echo = e
}

func (s *Server) InitDatabase(url string) {
	// Establish a connection to the PostgreSQL database using the provided DB_URL environment variable.
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("DB connection success!")

	// Automatically migrate the UserPostgres and Note models to the database.
	db.AutoMigrate(&models.UserPostgres{}, &models.Note{})
	s.db = db
}

func (s *Server) InitRedis(port string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:" + port,
		Password: "", // Password should be set if Redis requires authentication.
		DB:       0,  // Use default DB.
	})

	// Check if Redis connection established
	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal(err)
	}
	s.rdb = rdb
}

func (s *Server) InitGRPC(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen for RPC: %v", err)
	}

	grpc_server := grpc.NewServer()
	pb.RegisterTokenServiceServer(grpc_server, userrpc.NewRPCServer(s.rdb))

	log.Printf("RPC server listening at %v", lis.Addr())
	if err := grpc_server.Serve(lis); err != nil {
		log.Fatalf("RPC server failed to serve: %v", err)
	}
	s.grpc_server = grpc_server
}
