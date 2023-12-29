package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/billykore/todolist/internal/config"
	"github.com/billykore/todolist/internal/pkg/log"
	v1 "github.com/billykore/todolist/internal/proto/v1"
	"github.com/billykore/todolist/internal/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type HTTPServer struct {
	log *log.Logger
	cfg *config.Config
	svc *service.TodoService
}

func NewHTTPServer(log *log.Logger, cfg *config.Config, svc *service.TodoService) *HTTPServer {
	return &HTTPServer{
		log: log,
		cfg: cfg,
		svc: svc,
	}
}

func (hs *HTTPServer) Serve() {
	mux := runtime.NewServeMux()
	err := v1.RegisterTodoServiceHandlerServer(context.Background(), mux, hs.svc)
	if err != nil {
		hs.log.Fatalf("failed to register gateway: %v", err)
	}

	port := hs.cfg.HTTPPort
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	hs.log.Infof("Serving gRPC-Gateway for REST on :%s", port)
	if err = srv.ListenAndServe(); err != nil {
		hs.log.Fatalf("failed to serve http at port %s: %v", port, err)
	}
}
