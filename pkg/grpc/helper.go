package grpc

import (
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
)

func Dial(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
