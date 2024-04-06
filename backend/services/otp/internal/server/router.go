package server

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/services/otp/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	cfg        *config.Config
	log        *log.Logger
	router     *echo.Echo
	otpHandler *handler.OtpHandler
}

func NewRouter(cfg *config.Config, log *log.Logger, router *echo.Echo, otpHandler *handler.OtpHandler) *Router {
	return &Router{cfg: cfg, log: log, router: router, otpHandler: otpHandler}
}

func (r *Router) Run() {
	r.setRoutes()
	r.useMiddlewares()
	r.run()
}

func (r *Router) setRoutes() {
	r.router.POST("/otp/send", r.otpHandler.SendOtp)
	r.router.POST("/otp/verify", r.otpHandler.VerifyOtp)
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
