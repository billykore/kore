package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/billykore/todolist/libs/config"
	"github.com/billykore/todolist/libs/pkg/log"
	v1 "github.com/billykore/todolist/libs/proto/v1"
	"github.com/billykore/todolist/services/todo/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type HTTPServer struct {
	log     *log.Logger
	cfg     *config.Config
	todoSvc *service.TodoService
}

func NewHTTPServer(log *log.Logger, cfg *config.Config, todoSvc *service.TodoService) *HTTPServer {
	return &HTTPServer{
		log:     log,
		cfg:     cfg,
		todoSvc: todoSvc,
	}
}

func (hs *HTTPServer) Serve() {
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := v1.RegisterTodoServiceHandlerServer(ctx, mux, hs.todoSvc)
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
