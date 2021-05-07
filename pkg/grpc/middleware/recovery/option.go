package recovery

import (
	"context"

	"github.com/hidayatullahap/evermos/pkg/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	defaultOptions = &options{
		handlerFunc: nil,
	}
)

type options struct {
	handlerFunc HandlerFuncContext
}

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	*optCopy = *defaultOptions
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

// Option is a options wrapper
type Option func(*options)

// WithRecoveryHandler customizes the function for recovering from a panic.
func WithRecoveryHandler(f HandlerFuncContext) Option {
	return func(o *options) {
		o.handlerFunc = f
	}
}

// Options make recovery handler with Options
func Options() []Option {
	return []Option{
		WithRecoveryHandler(func(ctx context.Context, p interface{}) (err error) {
			return status.Errorf(codes.InternalError, "panic triggered: %v", p)
		}),
	}
}
