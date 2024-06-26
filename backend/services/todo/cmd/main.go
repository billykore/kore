package main

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/services/todo/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

type app struct {
	hs *server.HTTPServer
}

func newApp(hs *server.HTTPServer) *app {
	return &app{
		hs: hs,
	}
}

func main() {
	cfg := config.Get()
	todo := todoApp(cfg)
	todo.hs.Serve()
}
