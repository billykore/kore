//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/backend/internal/domain"
	"github.com/billykore/kore/backend/internal/infra/email"
	"github.com/billykore/kore/backend/internal/infra/http"
	"github.com/billykore/kore/backend/internal/infra/messaging"
	"github.com/billykore/kore/backend/internal/infra/persistence"
	"github.com/billykore/kore/backend/internal/infra/storage"
	"github.com/billykore/kore/backend/pkg"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func initApp(cfg *config.Config) *app {
	wire.Build(
		domain.ProviderSet,
		http.ProviderSet,
		email.ProviderSet,
		messaging.ProviderSet,
		persistence.ProviderSet,
		storage.ProviderSet,
		pkg.ProviderSet,
		echo.New,
		newApp,
	)
	return &app{}
}
