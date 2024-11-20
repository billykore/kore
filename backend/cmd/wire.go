//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/backend/domain"
	"github.com/billykore/kore/backend/infra/email"
	"github.com/billykore/kore/backend/infra/http"
	"github.com/billykore/kore/backend/infra/messaging"
	"github.com/billykore/kore/backend/infra/payment"
	"github.com/billykore/kore/backend/infra/shipment"
	"github.com/billykore/kore/backend/infra/storage"
	"github.com/billykore/kore/backend/pkg"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func initApp(cfg *config.Config) *app {
	wire.Build(
		domain.ProviderSet,
		storage.RepoProviderSet,
		storage.CacheProviderSet,
		email.ProviderSet,
		http.ProviderSet,
		messaging.ProviderSet,
		payment.ProviderSet,
		shipment.ProviderSet,
		pkg.ProviderSet,
		echo.New,
		newApp,
	)
	return &app{}
}
