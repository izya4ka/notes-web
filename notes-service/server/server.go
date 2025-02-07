package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/handlers"
	"github.com/izya4ka/notes-web/notes-service/middleware"
	"github.com/izya4ka/notes-web/notes-service/models"
	pb "github.com/izya4ka/notes-web/notes-service/proto"
	"github.com/izya4ka/notes-web/notes-service/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	db          *gorm.DB
	grpc_server *grpc.ClientConn
}

func (s *Server) InitDB(db_url string) {
	// Establish a connection to the PostgreSQL database using the provided DB_URL environment variable.
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Automatically migrate the UserPostgres and Note models to the database.
	db.AutoMigrate(&models.UserPostgres{}, &models.Note{})
	log.Print("DB connection success!")
	s.db = db
}

func (s *Server) InitGIN(port string) {
	router := gin.Default()

	token_service_client := pb.NewTokenServiceClient(s.grpc_server)

	router.GET("/notes", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			util.SendErrorResponse(c, err)
			return
		}
		handlers.GetNotes(c, s.db, username)
	})

	router.POST("/notes", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			util.SendErrorResponse(c, err)
			return
		}
		handlers.PostNotes(c, s.db, username)
	})

	router.GET("/notes/:id", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			util.SendErrorResponse(c, err)
			return
		}
		handlers.GetNote(c, s.db, username)
	})

	router.PUT("/notes/:id", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			util.SendErrorResponse(c, err)
			return
		}
		handlers.PutNode(c, s.db, username)
	})

	router.DELETE("/notes/:id", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			util.SendErrorResponse(c, err)
			return
		}
		handlers.DeleteNote(c, s.db, username)
	})

	router.Run("0.0.0.0:" + port)
}

func (s *Server) InitGRPC(port string) {
	conn, err := grpc.NewClient("user-service:"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("gRPC client failed to connect: %v", err)
	}
	s.grpc_server = conn
	log.Println("GRPC Client started!")
}
