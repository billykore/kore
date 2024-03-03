package server

import (
	"github.com/billykore/kore/pkg/config"
	"github.com/billykore/kore/pkg/log"
	"github.com/billykore/kore/services/todo/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	cfg     *config.Config
	log     *log.Logger
	router  *echo.Echo
	todoSvc *service.TodoService
}

func NewRouter(cfg *config.Config, log *log.Logger, router *echo.Echo, todoSvc *service.TodoService) *Router {
	return &Router{cfg: cfg, log: log, router: router, todoSvc: todoSvc}
}

func (r *Router) Run() {
	r.setRoutes()
	r.useMiddlewares()
	r.run()
}

func (r *Router) setRoutes() {
	tr := r.router.Group("/todos")
	tr.GET("", r.todoSvc.GetTodos)
	tr.GET("/:id", r.todoSvc.GetTodo)
	tr.POST("", r.todoSvc.AddTodo)
	tr.PUT("", r.todoSvc.SetDoneTodo)
	tr.DELETE("/:id", r.todoSvc.DeleteTodo)
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
