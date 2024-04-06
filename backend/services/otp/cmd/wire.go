//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/services/otp/internal/handler"
	"github.com/billykore/kore/backend/services/otp/internal/repo"
	"github.com/billykore/kore/backend/services/otp/internal/server"
	"github.com/billykore/kore/backend/services/otp/internal/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func otpApp(cfg *config.Config) *app {
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
