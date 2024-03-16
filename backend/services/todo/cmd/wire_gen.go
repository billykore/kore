// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/db"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/services/todo/repo"
	"github.com/billykore/kore/backend/services/todo/server"
	"github.com/billykore/kore/backend/services/todo/service"
	"github.com/billykore/kore/backend/services/todo/usecase"
	"github.com/labstack/echo/v4"
)

import (
	_ "github.com/joho/godotenv/autoload"
)

// Injectors from wire.go:

func todoApp(cfg *config.Config) *app {
	logger := log.NewLogger()
	echoEcho := echo.New()
	client := db.New(cfg)
	todoRepository := repo.NewTodoRepository(client)
	todoUsecase := usecase.NewTodoUsecase(logger, todoRepository)
	todoService := service.NewTodoService(todoUsecase)
	router := server.NewRouter(cfg, logger, echoEcho, todoService)
	httpServer := server.NewHTTPServer(router)
	mainApp := newApp(httpServer)
	return mainApp
}