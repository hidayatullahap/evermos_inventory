package recovery

import (
	"context"

	"github.com/hidayatullahap/evermos/pkg/grpc/codes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// HandlerFunc is a function that recovers from the panic `p` by returning an `error`.
type HandlerFunc func(p interface{}) (err error)

// HandlerFuncContext is a function that recovers from the panic `p` by returning an `error`.
// The context can be used to extract request scoped metadata and context values.
type HandlerFuncContext func(ctx context.Context, p interface{}) (err error)

// UnaryServerInterceptor returns a new unary server interceptor for panic recovery.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = recoverFrom(ctx, r, o.handlerFunc)
			}
		}()

		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor for panic recovery.
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := evaluateOptions(opts)
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = recoverFrom(stream.Context(), r, o.handlerFunc)
			}
		}()

		return handler(srv, stream)
	}
}

func recoverFrom(ctx context.Context, p interface{}, r HandlerFuncContext) error {
	if r == nil {
		return status.Errorf(codes.InternalError, "%s", p)
	}
	return r(ctx, p)
}
