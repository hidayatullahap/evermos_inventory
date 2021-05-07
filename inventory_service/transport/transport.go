package transport

import (
	"github.com/hidayatullahap/evermos/inventory_service/entity"
	"github.com/hidayatullahap/evermos/inventory_service/transport/grpc"
)

type Transport struct {
	GrpcServer *grpc.Server
}

func NewTransport(app *entity.App) *Transport {
	return &Transport{
		grpc.NewGrpcServer(app),
	}
}
