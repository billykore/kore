//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/todolist/libs/config"
	"github.com/billykore/todolist/libs/database"
	"github.com/billykore/todolist/libs/pkg"
	"github.com/billykore/todolist/services/todo/repository"
	"github.com/billykore/todolist/services/todo/server"
	"github.com/billykore/todolist/services/todo/service"
	"github.com/billykore/todolist/services/todo/usecase"
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
