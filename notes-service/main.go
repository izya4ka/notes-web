package main

import (
	"log"
	"net/http"
	"os"

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

func main() {
	router := gin.Default()

	// Establish a connection to the PostgreSQL database using the provided DB_URL environment variable.
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("DB connection success!")

	// Automatically migrate the UserPostgres and Note models to the database.
	db.AutoMigrate(&models.UserPostgres{}, &models.Note{})

	conn, err := grpc.NewClient("user-service:"+os.Getenv("GRPC_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("gRPC client failed to connect: %v", err)
	}
	defer conn.Close()

	token_service_client := pb.NewTokenServiceClient(conn)

	router.GET("/notes", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			util.SendErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token!")
			return
		}
		handlers.GetNotes(c, db, username)
	})

	router.POST("/notes", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			util.SendErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token!")
			return
		}
		handlers.PostNotes(c, db, username)
	})

	router.GET("/notes/:id", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			util.SendErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token!")
			return
		}
		handlers.GetNote(c, db, username)
	})

	router.PUT("/notes/:id", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			util.SendErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token!")
			return
		}
		handlers.PutNode(c, db, username)
	})

	router.DELETE("/notes/:id", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			util.SendErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token!")
			return
		}
		handlers.DeleteNote(c, db, username)
	})

	router.Run("0.0.0.0:" + os.Getenv("NOTES_SERVICE_PORT"))
}
