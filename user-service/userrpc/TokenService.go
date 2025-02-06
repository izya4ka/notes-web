package userrpc

import (
	"context"
	"errors"
	"log"

	"github.com/izya4ka/notes-web/user-service/database"
	pb "github.com/izya4ka/notes-web/user-service/proto"
	"github.com/izya4ka/notes-web/user-service/usererrors"
	"github.com/izya4ka/notes-web/user-service/util"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPC_server struct {
	pb.UnimplementedTokenServiceServer
	rdb *redis.Client
}

func (s *GRPC_server) GetUsername(ctx context.Context, req *pb.TokenRequest) (*pb.UsernameResponse, error) {
	raw_token := req.GetInput()

	token, err := util.UnrawToken(raw_token)

	if token == "" || err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid token")
	}

	username, err := database.GetUsername(ctx, s.rdb, token)
	if err != nil {
		if errors.Is(err, usererrors.ErrTimedOut) {
			return nil, status.Errorf(codes.DeadlineExceeded, "timed out")
		} else if errors.Is(err, usererrors.ErrInvalidToken) {
			return nil, status.Errorf(codes.NotFound, "invalid token")
		}
		log.Println("Error: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pb.UsernameResponse{
		Output: username,
	}, nil
}

func NewRPCServer(rdb *redis.Client) *GRPC_server {
	return &GRPC_server{rdb: rdb}
}
