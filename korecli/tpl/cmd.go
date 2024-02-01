package tpl

func MainTemplate() []byte {
	return []byte(`package main

import (
	"{{ .Mod }}/libs/config"
	"{{ .Mod }}/services/{{ .ServiceName }}/server"
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
	{{ .ServiceName }} := {{ .ServiceName }}App(cfg)

	go {{ .ServiceName }}.gs.Serve()
	{{ .ServiceName }}.hs.Serve()
}
`)
}

func WireTemplate() []byte {
	return []byte(`//go:build wireinject
// +build wireinject

package main

import (
	"{{ .Mod }}/libs/config"
	"{{ .Mod }}/libs/pkg"
	"{{ .Mod }}/services/{{ .ServiceName }}/repository"
	"{{ .Mod }}/services/{{ .ServiceName }}/server"
	"{{ .Mod }}/services/{{ .ServiceName }}/service"
	"{{ .Mod }}/services/{{ .ServiceName }}/usecase"
	"github.com/google/wire"
)

func {{ .ServiceName }}App(cfg *config.Config) *app {
	wire.Build(
		pkg.ProviderSet,
		repository.ProviderSet,
		usecase.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		newApp,
	)
	return &app{}
}
`)
}
