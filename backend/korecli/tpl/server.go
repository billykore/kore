package tpl

func ServerProviderTemplate() []byte {
	return []byte(`package server

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewRouter,
	NewHTTPServer,
)
`)
}

func RouterTemplate() []byte {
	return []byte(`package server

import (
	"{{ .GoMod }}/pkg/config"
	"{{ .GoMod }}/pkg/log"
	"{{ .GoMod }}/services/{{ .ServiceName }}/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	cfg     *config.Config
	log     *log.Logger
	router  *echo.Echo
	{{ .ServiceName }}Svc *service.{{ .StructName }}Service
}

func NewRouter(cfg *config.Config, log *log.Logger, router *echo.Echo, {{ .ServiceName }}Svc *service.{{ .StructName }}Service) *Router {
	return &Router{cfg: cfg, log: log, router: router, {{ .ServiceName }}Svc: {{ .ServiceName }}Svc}
}

func (r *Router) Run() {
	r.setRoutes()
	r.useMiddlewares()
	r.run()
}

func (r *Router) setRoutes() {
	r.router.GET("/greet", r.{{ .ServiceName }}Svc.Greet)
}

func (r *Router) useMiddlewares() {
	r.router.Use(middleware.Logger())
	r.router.Use(middleware.Recover())
}

func (r *Router) run() {
	port := r.cfg.HTTPPort
	if port == "" {
		port = "8080"
	}
	r.log.Infof("running on port [::%v]", port)
	if err := r.router.Start(":" + port); err != nil {
		r.log.Fatalf("failed to run on port [::%v]", port)
	}
}
`)
}

func HTTPServerTemplate() []byte {
	return []byte(`package server

type HTTPServer struct {
	router *Router
}

func NewHTTPServer(router *Router) *HTTPServer {
	return &HTTPServer{
		router: router,
	}
}

func (hs *HTTPServer) Serve() {
	hs.router.Run()
}
`)
}
