package main

import (
	"github.com/billykore/kore/backend/internal/infra/http/server"
	"github.com/billykore/kore/backend/internal/infra/messaging"
	"github.com/billykore/kore/backend/pkg/config"
	_ "github.com/joho/godotenv/autoload"
)

type app struct {
	ss *server.Server
	rc *messaging.Consumer
}

func newApp(ss *server.Server, rc *messaging.Consumer) *app {
	return &app{
		ss: ss,
		rc: rc,
	}
}

func main() {
	c := config.Get()
	a := initApp(c)
	go a.rc.Consume()
	a.ss.Serve()
}
