package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/middleware"
	pb "github.com/izya4ka/notes-web/notes-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	router := gin.Default()

	conn, err := grpc.NewClient("user-service:5002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("gRPC client failed to connect: %v", err)
	}
	defer conn.Close()

	token_service_client := pb.NewTokenServiceClient(conn)

	router.GET("/notes", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			c.String(http.StatusUnauthorized, "unauthorized")
			return
		}
		c.String(http.StatusOK, "notes and username: "+username)
	})

	router.POST("/notes", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			c.String(http.StatusUnauthorized, "unauthorized")
			return
		}
		c.String(http.StatusOK, "notes post"+" and username: "+username)
	})

	router.GET("/notes/:id", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			c.String(http.StatusUnauthorized, "unauthorized")
			return
		}
		c.String(http.StatusOK, "notes "+c.Param("id")+" get"+" and username: "+username)
	})

	router.PATCH("/notes/:id", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			c.String(http.StatusUnauthorized, "unauthorized")
			return
		}
		c.String(http.StatusOK, "notes"+c.Param("id")+" patch"+" and username: "+username)
	})

	router.DELETE("/notes/:id", func(c *gin.Context) {
		username, err := middleware.Auth(c, &token_service_client)
		if err != nil {
			c.String(http.StatusUnauthorized, "unauthorized")
			return
		}
		c.String(http.StatusOK, "notes"+c.Param("id")+" delete"+" and username: "+username)
	})

	router.Run("0.0.0.0:" + os.Getenv("NOTES_SERVICE_PORT"))
}
