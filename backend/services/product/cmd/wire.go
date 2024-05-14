//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/services/product/handler"
	"github.com/billykore/kore/backend/services/product/repo"
	"github.com/billykore/kore/backend/services/product/server"
	"github.com/billykore/kore/backend/services/product/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func productApp(cfg *config.Config) *app {
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
