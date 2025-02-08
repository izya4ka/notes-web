package main

import (
	"os"

	"github.com/izya4ka/notes-web/user-service/server"
)

// main initializes the Echo server, connects to the PostgreSQL database,
// and sets up the Redis client. If there is an errors in initialization, it will exit
func main() {
	var server server.Server

	go server.InitDatabase(os.Getenv("DB_URL"))
	go server.InitRedis(os.Getenv("REDIS_PORT"))
	go server.InitGRPC(os.Getenv("GRPC_PORT"))
	go server.InitEcho(os.Getenv("USER_SERVICE_PORT"))

	server.Shutdown()
}
