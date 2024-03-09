//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/db"
	"github.com/billykore/kore/backend/services/auth/repo"
	"github.com/billykore/kore/backend/services/auth/server"
	"github.com/billykore/kore/backend/services/auth/service"
	"github.com/billykore/kore/backend/services/auth/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func authApp(cfg *config.Config) *app {
	wire.Build(
		pkg.ProviderSet,
		db.ProviderSet,
		repo.ProviderSet,
		usecase.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		echo.New,
		newApp,
	)
	return &app{}
}
