//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/kore/libs/config"
	"github.com/billykore/kore/libs/pkg"
	"github.com/billykore/kore/services/user/repository"
	"github.com/billykore/kore/services/user/server"
	"github.com/billykore/kore/services/user/service"
	"github.com/billykore/kore/services/user/usecase"
	"github.com/google/wire"
)

func userApp(cfg *config.Config) *app {
	wire.Build(
		pkg.ProviderSet,
		repository.ProviderSet,
		usecase.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		newApp,
	)
	return &app{}
}
