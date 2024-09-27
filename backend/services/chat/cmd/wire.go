//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/services/chat/internal/handler"
	"github.com/billykore/kore/backend/services/chat/internal/server"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func chatApp(cfg *config.Config) *app {
	wire.Build(
		pkg.ProviderSet,
		handler.ProviderSet,
		server.ProviderSet,
		echo.New,
		newApp,
	)
	return &app{}
}
