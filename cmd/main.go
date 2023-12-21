package main

import (
	"github.com/billykore/todolist/internal/config"
	"github.com/billykore/todolist/internal/server"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type app struct {
	hs *server.HTTPServer
}

func newApp(hs *server.HTTPServer) *app {
	return &app{hs: hs}
}

func setEnv(env string) {
	if env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	cfg := config.Get()
	setEnv(cfg.Env)
	todo := todoApp(cfg)
	todo.hs.Serve()
}
