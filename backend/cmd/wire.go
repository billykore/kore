//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/backend/internal/app"
	"github.com/billykore/kore/backend/internal/infra"
	"github.com/billykore/kore/backend/internal/infra/database/postgres"
	"github.com/billykore/kore/backend/internal/infra/http/handler"
	"github.com/billykore/kore/backend/pkg"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func initKore(cfg *config.Config) *kore {
	wire.Build(
		app.ProviderSet,
		infra.ProviderSet,
		handler.ProviderSet,
		postgres.ProviderSet,
		pkg.ProviderSet,
		echo.New,
		newKore,
	)
	return &kore{}
}
