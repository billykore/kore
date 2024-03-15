//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/db"
	"github.com/billykore/kore/backend/services/todo/repo"
	"github.com/billykore/kore/backend/services/todo/server"
	"github.com/billykore/kore/backend/services/todo/service"
	"github.com/billykore/kore/backend/services/todo/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func todoApp(cfg *config.Config) *app {
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
