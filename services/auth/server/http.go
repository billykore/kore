package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/billykore/todolist/libs/config"
	"github.com/billykore/todolist/libs/pkg/log"
	"github.com/billykore/todolist/libs/proto/v1"
	"github.com/billykore/todolist/services/auth/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type HTTPServer struct {
	log     *log.Logger
	cfg     *config.Config
	authSvc *service.AuthService
}

func NewHTTPServer(log *log.Logger, cfg *config.Config, authSvc *service.AuthService) *HTTPServer {
	return &HTTPServer{
		log:     log,
		cfg:     cfg,
		authSvc: authSvc,
	}
}

func (hs *HTTPServer) Serve() {
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := v1.RegisterAuthorizationHandlerServer(ctx, mux, hs.authSvc)
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
