package main

import (
	"os"

	"github.com/izya4ka/notes-web/notes-service/server"
)

func main() {
	var server server.Server

	go server.InitDB(os.Getenv("DB_URL"))
	go server.InitGIN(os.Getenv("NOTES_SERVICE_PORT"))

	server.Shutdown()
}
