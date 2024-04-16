//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/backend/libs/config"
	"github.com/billykore/kore/backend/libs/pkg"
	"github.com/billykore/kore/backend/services/product/repo"
	"github.com/billykore/kore/backend/services/product/server"
	"github.com/billykore/kore/backend/services/product/service"
	"github.com/billykore/kore/backend/services/product/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func productApp(cfg *config.Config) *app {
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
