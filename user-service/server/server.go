package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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
	mutex       sync.Mutex
}

func (s *Server) InitEcho(port string) {
	s.mutex.Lock()
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
	s.echo = e
	s.mutex.Unlock()
	e.Start(":" + port)
}

func (s *Server) InitDatabase(url string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
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
	s.mutex.Lock()
	defer s.mutex.Unlock()
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
	s.mutex.Lock()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen for RPC: %v", err)
	}

	grpc_server := grpc.NewServer()
	s.grpc_server = grpc_server
	pb.RegisterTokenServiceServer(grpc_server, userrpc.NewRPCServer(s.rdb))

	log.Printf("RPC server listening at %v", lis.Addr())
	s.mutex.Unlock()
	if err := grpc_server.Serve(lis); err != nil {
		log.Fatalf("RPC server failed to serve: %v", err)
	}
}

func (s *Server) Shutdown() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Starting graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.echo.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Echo stopping error: %v", err)
	}
	log.Println("Echo stopped!")

	sqldb, serr := s.db.DB()
	if serr == nil {
		if err := sqldb.Close(); err != nil {
			log.Fatalf("Database stopping error: %v", err)
		}
	} else {
		log.Fatalf("Database stopping error: %v", serr)
	}
	log.Println("Database connection stopped!")

	if err := s.rdb.Close(); err != nil {
		log.Fatalf("Redis client stopping error: %v", err)
	}
	log.Println("Redis client stopped!")

	s.grpc_server.GracefulStop()
	log.Println("GRPC server stopped!")
	os.Exit(0)
}
