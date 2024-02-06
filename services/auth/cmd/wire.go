//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/libs/config"
	"github.com/billykore/kore/libs/database"
	"github.com/billykore/kore/libs/pkg"
	"github.com/billykore/kore/services/auth/repository"
	"github.com/billykore/kore/services/auth/server"
	"github.com/billykore/kore/services/auth/service"
	"github.com/billykore/kore/services/auth/usecase"
	"github.com/google/wire"
)

func authApp(cfg *config.Config) *app {
	wire.Build(
		pkg.ProviderSet,
		database.ProviderSet,
		repository.ProviderSet,
		usecase.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		newApp,
	)
	return &app{}
}
