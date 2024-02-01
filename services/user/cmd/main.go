package main

import (
	"github.com/billykore/kore/libs/config"
	"github.com/billykore/kore/services/user/server"
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
	user := userApp(cfg)

	go user.gs.Serve()
	user.hs.Serve()
}
