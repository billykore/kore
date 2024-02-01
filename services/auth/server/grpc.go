package server

import (
	"fmt"
	"net"

	"github.com/billykore/kore/libs/config"
	"github.com/billykore/kore/libs/pkg/log"
	"github.com/billykore/kore/libs/proto/v1"
	"github.com/billykore/kore/services/auth/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	log     *log.Logger
	cfg     *config.Config
	authSvc *service.AuthService
}

func NewGRPCServer(log *log.Logger, cfg *config.Config, authSvc *service.AuthService) *GRPCServer {
	return &GRPCServer{
		log:     log,
		cfg:     cfg,
		authSvc: authSvc,
	}
}

func (gs *GRPCServer) Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gs.cfg.GRPCPort))
	if err != nil {
		gs.log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	v1.RegisterAuthorizationServer(srv, gs.authSvc)
	reflection.Register(srv)

	gs.log.Infof("Run on grpc server port %s", gs.cfg.GRPCPort)
	if err = srv.Serve(lis); err != nil {
		gs.log.Fatalf("failed to serve grpc: %v", err)
	}
}
