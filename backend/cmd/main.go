package main

import (
	"github.com/billykore/kore/backend/internal/infra/http"
	"github.com/billykore/kore/backend/pkg/config"
	_ "github.com/joho/godotenv/autoload"
)

type kore struct {
	hs *http.Server
}

func newKore(hs *http.Server) *kore {
	return &kore{
		hs: hs,
	}
}

func main() {
	c := config.Get()
	k := initKore(c)
	k.hs.Serve()
}
