package server

import (
	"github.com/billykore/todolist/internal/config"
)

type HTTPServer struct {
	config *config.Config
	router *Router
}

func NewHTTPServer(cfg *config.Config, router *Router) *HTTPServer {
	return &HTTPServer{
		config: cfg,
		router: router,
	}
}

func (hs *HTTPServer) Serve() {
	hs.router.Run(hs.config.Port)
}
