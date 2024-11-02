package main

import (
	"github.com/billykore/kore/backend/internal/infra/http"
	"github.com/billykore/kore/backend/internal/infra/messaging"
	"github.com/billykore/kore/backend/pkg/config"
	_ "github.com/joho/godotenv/autoload"
)

type kore struct {
	hs *http.Server
	cs *messaging.Consumer
}

func newKore(hs *http.Server, cs *messaging.Consumer) *kore {
	return &kore{
		hs: hs,
		cs: cs,
	}
}

func main() {
	c := config.Get()
	k := initKore(c)
	go k.cs.Consume()
	k.hs.Serve()
}
