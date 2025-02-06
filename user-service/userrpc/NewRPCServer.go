package userrpc

import (
	"github.com/redis/go-redis/v9"
)

func NewRPCServer(rdb *redis.Client) *GRPC_server {
	return &GRPC_server{rdb: rdb}
}
