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
}

func newApp(hs *server.HTTPServer) *app {
	return &app{
		hs: hs,
	}
}

func main() {
	cfg := config.Get()
	{{ .ServiceName }} := {{ .ServiceName }}App(cfg)
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
	"{{ .Mod }}/services/{{ .ServiceName }}/repo"
	"{{ .Mod }}/services/{{ .ServiceName }}/server"
	"{{ .Mod }}/services/{{ .ServiceName }}/service"
	"{{ .Mod }}/services/{{ .ServiceName }}/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func {{ .ServiceName }}App(cfg *config.Config) *app {
	wire.Build(
		pkg.ProviderSet,
		repo.ProviderSet,
		usecase.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		echo.New,
		newApp,
	)
	return &app{}
}
`)
}
