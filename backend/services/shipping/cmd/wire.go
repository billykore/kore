//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/services/shipping/handler"
	"github.com/billykore/kore/backend/services/shipping/repo"
	"github.com/billykore/kore/backend/services/shipping/server"
	"github.com/billykore/kore/backend/services/shipping/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func shippingApp(cfg *config.Config) *app {
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
