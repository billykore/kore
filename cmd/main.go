package main

import (
	"github.com/billykore/todolist/internal/config"
	"github.com/billykore/todolist/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

type app struct {
	hs *server.HTTPServer
	gs *server.GRPCServer
}

func newApp(hs *server.HTTPServer, gs *server.GRPCServer) *app {
	return &app{
		hs: hs,
		gs: gs,
	}
}

func main() {
	cfg := config.Get()
	todo := todoApp(cfg)

	go todo.gs.Serve()
	todo.hs.Serve()
}
