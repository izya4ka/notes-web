package main

import (
	"log"
	"os"

	"github.com/izya4ka/notes-web/gateway/middleware"
	pb "github.com/izya4ka/notes-web/gateway/proto"
	"github.com/izya4ka/notes-web/gateway/util"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger)

	if os.Getenv("DEBUG") != "" {
		e.Debug = true
	}

	userServiceURL := os.Getenv("USER_SERVICE_URL")
	notesServiceURL := os.Getenv("NOTES_SERVICE_URL")
	grpcAddress := os.Getenv("GRPC_ADDRESS")
	gatewayPort := os.Getenv("GATEWAY_PORT")

	conn, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		util.LogFatalf("gRPC didn't start: %s", err)
	}

	rpc := pb.NewTokenServiceClient(conn)
	util.LogInfof("gRPC client started!")

	api := e.Group("/api")
	v1 := api.Group("/v1")

	v1.Any("/notes*", func(c echo.Context) error {
		return Handler(c, notesServiceURL)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return middleware.Auth(next, &rpc)
	})

	v1.Any("/user*", func(c echo.Context) error {
		return Handler(c, userServiceURL)
	})

	log.Fatal(e.Start(":" + gatewayPort))
}
