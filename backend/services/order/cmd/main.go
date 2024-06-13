package main

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/services/order/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

type app struct {
	hs *server.HTTPServer
	bs *server.BrokerServer
}

func newApp(hs *server.HTTPServer, bs *server.BrokerServer) *app {
	return &app{
		hs: hs,
		bs: bs,
	}
}

func main() {
	cfg := config.Get()
	order := orderApp(cfg)
	go order.bs.Serve()
	order.hs.Serve()
}
