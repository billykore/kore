package main

import (
	"github.com/billykore/kore/backend/internal/infra/http"
	"github.com/billykore/kore/backend/internal/infra/messaging"
	"github.com/billykore/kore/backend/pkg/config"
	_ "github.com/joho/godotenv/autoload"
)

type app struct {
	hs *http.Server
	mc *messaging.Consumer
}

func newApp(hs *http.Server, mc *messaging.Consumer) *app {
	return &app{
		hs: hs,
		mc: mc,
	}
}

func main() {
	c := config.Get()
	a := initApp(c)
	go a.mc.Consume()
	a.hs.Serve()
}
