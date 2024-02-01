//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/libs/config"
	"github.com/billykore/kore/libs/database"
	"github.com/billykore/kore/libs/pkg"
	"github.com/billykore/kore/services/todo/repository"
	"github.com/billykore/kore/services/todo/server"
	"github.com/billykore/kore/services/todo/service"
	"github.com/billykore/kore/services/todo/usecase"
	"github.com/google/wire"
)

func todoApp(cfg *config.Config) *app {
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
