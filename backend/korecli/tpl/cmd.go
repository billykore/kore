package tpl

func MainTemplate() []byte {
	return []byte(`package main

import (
	"{{ .GoMod }}/pkg/config"
	"{{ .GoMod }}/services/{{ .ServiceName }}/server"
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
	"{{ .GoMod }}/libs/pkg/config"
	"{{ .GoMod }}/libs/pkg"
	"{{ .GoMod }}/services/{{ .ServiceName }}/repo"
	"{{ .GoMod }}/services/{{ .ServiceName }}/server"
	"{{ .GoMod }}/services/{{ .ServiceName }}/service"
	"{{ .GoMod }}/services/{{ .ServiceName }}/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func {{ .ServiceName }}App(cfg *config.Config) *app {
	wire.Build(
		pkg.ProviderSet,
		repo.ProviderSet,
		usecase.ProviderSet,
		handler.ProviderSet,
		server.ProviderSet,
		echo.New,
		newApp,
	)
	return &app{}
}
`)
}
