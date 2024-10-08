// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/net/websocket"
	"github.com/billykore/kore/backend/services/chat/internal/handler"
	"github.com/billykore/kore/backend/services/chat/internal/server"
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
	chatHandler := handler.NewChatHandler(pool)
	router := server.NewRouter(cfg, logger, echoEcho, pool, chatHandler)
	httpServer := server.NewHTTPServer(router)
	mainApp := newApp(httpServer)
	return mainApp
}
