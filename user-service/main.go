package main

import (
	"os"

	"github.com/izya4ka/notes-web/user-service/server"
)

// main initializes the Echo server, connects to the PostgreSQL database,
// and sets up the Redis client. If there is an errors in initialization, it will exit
func main() {
	var s server.Server

	s.InitDatabase(os.Getenv("DB_URL"))
	s.InitRedis(os.Getenv("REDIS_PORT"))
	go s.InitGRPC(os.Getenv("GRPC_PORT"))
	s.InitEcho(os.Getenv("USER_SERVICE_PORT"))

}
