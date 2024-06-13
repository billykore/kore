// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/db"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/services/product/internal/handler"
	"github.com/billykore/kore/backend/services/product/internal/repo"
	"github.com/billykore/kore/backend/services/product/internal/server"
	"github.com/billykore/kore/backend/services/product/internal/usecase"
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
	productRepo := repo.NewProductRepository(gormDB)
	productUsecase := usecase.NewProductUsecase(logger, productRepo)
	productHandler := handler.NewProductHandler(productUsecase)
	productCategoryRepo := repo.NewProductCategoryRepository(gormDB)
	productCategoryUsecase := usecase.NewProductCategoryUsecase(logger, productCategoryRepo)
	productCategoryHandler := handler.NewProductCategoryHandler(productCategoryUsecase)
	discountRepo := repo.NewDiscountRepository(gormDB)
	discountUsecase := usecase.NewDiscountUsecase(discountRepo)
	discountHandler := handler.NewDiscountHandler(discountUsecase)
	cartRepo := repo.NewCartRepository(gormDB)
	cartUsecase := usecase.NewCartUsecase(logger, cartRepo)
	cartHandler := handler.NewCartHandler(cartUsecase)
	router := server.NewRouter(cfg, logger, echoEcho, productHandler, productCategoryHandler, discountHandler, cartHandler)
	httpServer := server.NewHTTPServer(router)
	mainApp := newApp(httpServer)
	return mainApp
}
