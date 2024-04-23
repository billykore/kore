// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/db"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/services/product/repo"
	"github.com/billykore/kore/backend/services/product/server"
	"github.com/billykore/kore/backend/services/product/service"
	"github.com/billykore/kore/backend/services/product/usecase"
	"github.com/labstack/echo/v4"
)

import (
	_ "github.com/joho/godotenv/autoload"
)

// Injectors from wire.go:

func productApp(cfg *config.Config) *app {
	logger := log.NewLogger()
	echoEcho := echo.New()
	gormDB := db.NewPostgres(cfg)
	productRepository := repo.NewProductRepository(gormDB)
	productUsecase := usecase.NewProductUsecase(logger, productRepository)
	productService := service.NewProductService(productUsecase)
	router := server.NewRouter(cfg, logger, echoEcho, productService)
	httpServer := server.NewHTTPServer(router)
	mainApp := newApp(httpServer)
	return mainApp
}
