package main

import (
	"github.com/billykore/kore/libs/config"
	"github.com/billykore/kore/services/auth/server"
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
	auth := authApp(cfg)

	go auth.gs.Serve()
	auth.hs.Serve()
}
