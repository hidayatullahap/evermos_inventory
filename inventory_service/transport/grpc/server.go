package grpc

import (
	"fmt"
	"log"
	"net"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/hidayatullahap/evermos/inventory_service/entity"
	"github.com/hidayatullahap/evermos/pkg/grpc/middleware/validator"
	"github.com/hidayatullahap/evermos/pkg/proto/inventory"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	app *entity.App
}

func (a *Server) Start() {
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		fmt.Println("gRPC server start at port :" + port)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
		validator.UnaryServerInterceptor(),
	)))

	hsrv := health.NewServer()
	hsrv.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, hsrv)

	inventory.RegisterInventoryServer(s, NewGrpcHandler(a.app))

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func NewGrpcServer(app *entity.App) *Server {
	return &Server{
		app: app,
	}
}
