//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/libs/config"
	"github.com/billykore/kore/libs/pkg"
	"github.com/billykore/kore/services/book/repo"
	"github.com/billykore/kore/services/book/server"
	"github.com/billykore/kore/services/book/service"
	"github.com/billykore/kore/services/book/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func bookApp(cfg *config.Config) *app {
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
