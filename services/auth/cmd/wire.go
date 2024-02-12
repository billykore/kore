//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/libs/config"
	"github.com/billykore/kore/libs/db"
	"github.com/billykore/kore/libs/pkg"
	"github.com/billykore/kore/services/auth/repo"
	"github.com/billykore/kore/services/auth/server"
	"github.com/billykore/kore/services/auth/service"
	"github.com/billykore/kore/services/auth/usecase"
	"github.com/google/wire"
)

func authApp(cfg *config.Config) *app {
	wire.Build(
		pkg.ProviderSet,
		db.ProviderSet,
		repo.ProviderSet,
		usecase.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		newApp,
	)
	return &app{}
}
