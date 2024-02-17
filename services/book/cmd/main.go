package main

import (
	"github.com/billykore/kore/libs/config"
	"github.com/billykore/kore/services/book/server"
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
	book := bookApp(cfg)
	book.hs.Serve()
}
