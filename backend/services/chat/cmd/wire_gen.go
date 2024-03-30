// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/websocket"
	"github.com/billykore/kore/backend/services/chat/repo"
	"github.com/billykore/kore/backend/services/chat/server"
	"github.com/billykore/kore/backend/services/chat/service"
	"github.com/billykore/kore/backend/services/chat/usecase"
	"github.com/labstack/echo/v4"
)

import (
	_ "github.com/joho/godotenv/autoload"
)

// Injectors from wire.go:

func chatApp(cfg *config.Config) *app {
	logger := log.NewLogger()
	echoEcho := echo.New()
	pool := websocket.NewPool()
	greeterRepository := repo.NewChatRepository()
	chatUsecase := usecase.NewChatUsecase(logger, greeterRepository)
	chatService := service.NewChatService(chatUsecase, pool)
	router := server.NewRouter(cfg, logger, echoEcho, pool, chatService)
	httpServer := server.NewHTTPServer(router)
	mainApp := newApp(httpServer)
	return mainApp
}
