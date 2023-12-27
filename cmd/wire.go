//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/todolist/internal/config"
	"github.com/billykore/todolist/internal/database"
	"github.com/billykore/todolist/internal/pkg"
	"github.com/billykore/todolist/internal/repository"
	"github.com/billykore/todolist/internal/server"
	"github.com/billykore/todolist/internal/service"
	"github.com/billykore/todolist/internal/usecase"
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
