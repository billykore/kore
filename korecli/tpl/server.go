package tpl

func ServerProviderTemplate() []byte {
	return []byte(`package server

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewGRPCServer,
	NewHTTPServer,
)
`)
}

func HTTPServerTemplate() []byte {
	return []byte(`package server

import (
	"context"
	"fmt"
	"net/http"

	"{{ .Mod }}/libs/config"
	"{{ .Mod }}/libs/pkg/log"
	"{{ .Mod }}/libs/proto/v1"
	"{{ .Mod }}/services/{{ .ServiceName }}/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type HTTPServer struct {
	log     *log.Logger
	cfg     *config.Config
	{{ .ServiceName }}Svc *service.{{ .StructName }}Service
}

func NewHTTPServer(log *log.Logger, cfg *config.Config, {{ .ServiceName }}Svc *service.{{ .StructName }}Service) *HTTPServer {
	return &HTTPServer{
		log:     log,
		cfg:     cfg,
		{{ .ServiceName }}Svc: {{ .ServiceName }}Svc,
	}
}

func (hs *HTTPServer) Serve() {
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := v1.RegisterGreeterHandlerServer(ctx, mux, hs.{{ .ServiceName }}Svc)
	if err != nil {
		hs.log.Fatalf("failed to register gateway: %v", err)
	}

	port := hs.cfg.HTTPPort
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	hs.log.Infof("Serving gRPC-Gateway for REST on port %s", port)
	if err = srv.ListenAndServe(); err != nil {
		hs.log.Fatalf("failed to serve http at port %s: %v", port, err)
	}
}
`)
}

func GRPCServerTemplate() []byte {
	return []byte(`package server

import (
	"fmt"
	"net"

	"{{ .Mod }}/libs/config"
	"{{ .Mod }}/libs/pkg/log"
	"{{ .Mod }}/libs/proto/v1"
	"{{ .Mod }}/services/{{ .ServiceName }}/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	log     *log.Logger
	cfg     *config.Config
	{{ .ServiceName }}Svc *service.{{ .StructName }}Service
}

func NewGRPCServer(log *log.Logger, cfg *config.Config, {{ .ServiceName }}Svc *service.{{ .StructName }}Service) *GRPCServer {
	return &GRPCServer{
		log:     log,
		cfg:     cfg,
		{{ .ServiceName }}Svc: {{ .ServiceName }}Svc,
	}
}

func (gs *GRPCServer) Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gs.cfg.GRPCPort))
	if err != nil {
		gs.log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	v1.RegisterGreeterServer(srv, gs.{{ .ServiceName }}Svc)
	reflection.Register(srv)

	gs.log.Infof("Run on grpc server port %s", gs.cfg.GRPCPort)
	if err = srv.Serve(lis); err != nil {
		gs.log.Fatalf("failed to serve grpc: %v", err)
	}
}

`)
}
