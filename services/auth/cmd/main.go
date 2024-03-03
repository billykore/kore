package main

import (
	"github.com/billykore/kore/pkg/config"
	"github.com/billykore/kore/services/auth/server"
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
	auth := authApp(cfg)
	auth.hs.Serve()
}
