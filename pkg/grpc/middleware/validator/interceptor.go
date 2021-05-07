package validator

import (
	"context"

	"github.com/hidayatullahap/evermos/pkg/grpc/codes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type validator interface {
	Validate() error
}

// UnaryServerInterceptor returns a new unary server interceptor that validates incoming messages.
//
// Invalid messages will be rejected with `InvalidArgument` before reaching any userspace handlers.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if v, ok := req.(validator); ok {
			if err := v.Validate(); err != nil {
				return nil, status.Errorf(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	}
}
