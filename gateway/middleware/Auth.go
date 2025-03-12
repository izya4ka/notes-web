package middleware

import (
	"context"
	"time"

	"github.com/izya4ka/notes-web/gateway/gateerrors"
	pb "github.com/izya4ka/notes-web/gateway/proto"
	"github.com/izya4ka/notes-web/gateway/util"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Auth(next echo.HandlerFunc, rpc *pb.TokenServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		util.LogDebugf("Request headers received: %v", c.Request().Header)
		username, err := (*rpc).GetUsername(ctx, &pb.TokenRequest{Input: c.Request().Header.Get("Authorization")})
		if err != nil {
			if st, ok := status.FromError(err); ok {
				switch st.Code() {
				case codes.InvalidArgument:
					return util.SendErrorResponse(c, gateerrors.ErrInvalidToken)
				case codes.DeadlineExceeded:
					return util.SendErrorResponse(c, gateerrors.ErrTimedOut)
				case codes.NotFound:
					return util.SendErrorResponse(c, gateerrors.ErrInvalidToken)
				default:
					util.LogErrorf("Error: %s", err)
					return util.SendErrorResponse(c, gateerrors.ErrInternal)
				}
			}
		}
		util.LogDebugf("Username received from user-service: %s", username.Output)
		c.Request().Header.Del("Authorization")
		c.Request().Header.Set("Username", username.Output)
		if err := next(c); err != nil {
			c.Error(err)
		}
		util.LogDebugf("Request headers to send to service: %v", c.Request().Header)
		return nil
	}
}
