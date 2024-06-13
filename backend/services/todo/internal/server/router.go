package server

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/services/todo/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	cfg         *config.Config
	log         *log.Logger
	router      *echo.Echo
	todoHandler *handler.TodoHandler
}

func NewRouter(cfg *config.Config, log *log.Logger, router *echo.Echo, todoHandler *handler.TodoHandler) *Router {
	return &Router{cfg: cfg, log: log, router: router, todoHandler: todoHandler}
}

func (r *Router) Run() {
	r.setRoutes()
	r.useMiddlewares()
	r.run()
}

func (r *Router) setRoutes() {
	tr := r.router.Group("/todos")
	tr.GET("", r.todoHandler.GetTodos)
	tr.GET("/:id", r.todoHandler.GetTodo)
	tr.POST("", r.todoHandler.AddTodo)
	tr.PUT("/:id", r.todoHandler.SetDoneTodo)
	tr.DELETE("/:id", r.todoHandler.DeleteTodo)
}

func (r *Router) useMiddlewares() {
	r.router.Use(middleware.Logger())
	r.router.Use(middleware.Recover())
	r.router.Use(middleware.CORS())
}

func (r *Router) run() {
	port := r.cfg.HTTPPort
	r.log.Infof("running on port [::%v]", port)
	if err := r.router.Start(":" + port); err != nil {
		r.log.Fatalf("failed to run on port [::%v]", port)
	}
}
