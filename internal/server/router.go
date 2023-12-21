package server

import (
	"log"

	"github.com/billykore/todolist/internal/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine      *gin.Engine
	todoHandler *handler.TodoHandler
}

func NewRouter(engine *gin.Engine, todoHandler *handler.TodoHandler) *Router {
	return &Router{
		engine:      engine,
		todoHandler: todoHandler,
	}
}

func (r *Router) Run(port string) {
	r.routes()
	r.run(port)
}

func (r *Router) routes() {
	r.todoRoutes()
}

func (r *Router) run(port string) {
	if err := r.engine.Run(":" + port); err != nil {
		log.Fatalf("failed to run on port %s: %v", port, err)
	}
}

func (r *Router) todoRoutes() {
	tr := r.engine.Group("/todos")
	tr.GET("", r.todoHandler.GetTodos)
	tr.POST("", r.todoHandler.AddTodo)
	tr.GET("/:id", r.todoHandler.GetTodo)
	tr.PUT("/:id", r.todoHandler.SetDoneTodo)
	tr.DELETE("/:id", r.todoHandler.DeleteTodo)
}
