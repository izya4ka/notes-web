package middleware

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/noteserrors"
	pb "github.com/izya4ka/notes-web/notes-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Auth(c *gin.Context, rpc *pb.TokenServiceClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	username, err := (*rpc).GetUsername(ctx, &pb.TokenRequest{Input: c.Request.Header.Get("Authorization")})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.InvalidArgument:
				return "", noteserrors.ErrInvalidToken
			case codes.DeadlineExceeded:
				return "", noteserrors.ErrTimedOut
			case codes.NotFound:
				return "", noteserrors.ErrInvalidToken
			default:
				log.Println("Error: ", st.Err())
				return "", noteserrors.ErrInternal
			}
		}
	}

	return username.Output, nil
}
