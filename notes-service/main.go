package main

import (
	"os"

	"github.com/izya4ka/notes-web/notes-service/server"
)

func main() {
	var server server.Server

	server.InitDB(os.Getenv("DB_URL"))
	server.InitGRPC(os.Getenv("GRPC_PORT"))
	server.InitGIN(os.Getenv("NOTES_SERVICE_PORT"))
}
