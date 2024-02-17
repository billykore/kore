package server

import (
	"github.com/billykore/kore/libs/config"
	"github.com/billykore/kore/libs/pkg/log"
	"github.com/billykore/kore/services/book/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	cfg     *config.Config
	log     *log.Logger
	router  *echo.Echo
	bookSvc *service.BookService
}

func NewRouter(cfg *config.Config, log *log.Logger, router *echo.Echo, bookSvc *service.BookService) *Router {
	return &Router{cfg: cfg, log: log, router: router, bookSvc: bookSvc}
}

func (r *Router) Run() {
	r.setRoutes()
	r.useMiddlewares()
	r.run()
}

func (r *Router) setRoutes() {
	r.router.GET("/greet", r.bookSvc.Greet)
}

func (r *Router) useMiddlewares() {
	r.router.Use(middleware.Logger())
}

func (r *Router) run() {
	port := r.cfg.HTTPPort
	if port == "" {
		port = "8080"
	}
	r.log.Infof("running on port ::[:%v]", port)
	if err := r.router.Start(":" + port); err != nil {
		r.log.Fatalf("failed to run on port [::%v]", port)
	}
}
