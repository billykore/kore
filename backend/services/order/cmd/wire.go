//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/services/order/handler"
	"github.com/billykore/kore/backend/services/order/repo"
	"github.com/billykore/kore/backend/services/order/server"
	"github.com/billykore/kore/backend/services/order/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func orderApp(cfg *config.Config) *app {
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
