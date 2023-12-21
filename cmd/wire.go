//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/todolist/internal/config"
	"github.com/billykore/todolist/internal/database"
	"github.com/billykore/todolist/internal/handler"
	"github.com/billykore/todolist/internal/pkg"
	"github.com/billykore/todolist/internal/repository"
	"github.com/billykore/todolist/internal/server"
	"github.com/billykore/todolist/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func todoApp(cfg *config.Config) *app {
	wire.Build(
		gin.Default,
		pkg.ProviderSet,
		database.ProviderSet,
		repository.ProviderSet,
		usecase.ProviderSet,
		handler.ProviderSet,
		server.ProviderSet,
		newApp,
	)
	return &app{}
}
