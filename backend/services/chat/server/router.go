package server

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/websocket"
	"github.com/billykore/kore/backend/services/chat/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	cfg     *config.Config
	log     *log.Logger
	router  *echo.Echo
	pool    *websocket.Pool
	chatSvc *handler.ChatHandler
}

func NewRouter(cfg *config.Config, log *log.Logger, router *echo.Echo, pool *websocket.Pool, chatSvc *handler.ChatHandler) *Router {
	return &Router{
		cfg:     cfg,
		log:     log,
		router:  router,
		pool:    pool,
		chatSvc: chatSvc,
	}
}

func (r *Router) Run() {
	r.setRoutes()
	r.useMiddlewares()

	go r.pool.Start()
	r.run()
}

func (r *Router) setRoutes() {
	r.router.GET("/chat", r.chatSvc.Chat)
}

func (r *Router) useMiddlewares() {
	r.router.Use(middleware.Logger())
	r.router.Use(middleware.Recover())
}

func (r *Router) run() {
	port := r.cfg.HTTPPort
	r.log.Infof("running on port [::%v]", port)
	if err := r.router.Start(":" + port); err != nil {
		r.log.Fatalf("failed to run on port [::%v]", port)
	}
}
