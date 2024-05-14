package server

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/services/order/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	cfg      *config.Config
	log      *log.Logger
	router   *echo.Echo
	orderSvc *handler.OrderHandler
}

func NewRouter(cfg *config.Config, log *log.Logger, router *echo.Echo, orderSvc *handler.OrderHandler) *Router {
	return &Router{cfg: cfg, log: log, router: router, orderSvc: orderSvc}
}

func (r *Router) Run() {
	r.setRoutes()
	r.useMiddlewares()
	r.run()
}

func (r *Router) setRoutes() {
	r.router.POST("/orders/checkout", r.orderSvc.Checkout)
	r.router.GET("/orders/:orderId", r.orderSvc.GetOrderById)
	r.router.POST("/orders/:orderId/payment", r.orderSvc.PayOrder)
	r.router.POST("/orders/:orderId/shipping", r.orderSvc.ShipOrder)
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
