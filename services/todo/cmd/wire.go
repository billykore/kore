//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/libs/config"
	"github.com/billykore/kore/libs/db"
	"github.com/billykore/kore/libs/pkg"
	"github.com/billykore/kore/services/todo/repo"
	"github.com/billykore/kore/services/todo/server"
	"github.com/billykore/kore/services/todo/service"
	"github.com/billykore/kore/services/todo/usecase"
	"github.com/google/wire"
)

func todoApp(cfg *config.Config) *app {
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
