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
	"github.com/billykore/kore/backend/services/product/handler"
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
	productService := handler.NewProductHandler(productUsecase)
	productCategoryRepository := repo.NewProductCategoryRepository(gormDB)
	productCategoryUsecase := usecase.NewProductCategoryUsecase(logger, productCategoryRepository)
	productCategoryService := handler.NewProductCategoryHandler(productCategoryUsecase)
	discountRepository := repo.NewDiscountRepository(gormDB)
	discountUsecase := usecase.NewDiscountUsecase(discountRepository)
	discountService := handler.NewDiscountHandler(discountUsecase)
	cartRepository := repo.NewCartRepository(gormDB)
	cartUsecase := usecase.NewCartUsecase(logger, cartRepository)
	cartService := handler.NewCartHandler(cartUsecase)
	router := server.NewRouter(cfg, logger, echoEcho, productService, productCategoryService, discountService, cartService)
	httpServer := server.NewHTTPServer(router)
	mainApp := newApp(httpServer)
	return mainApp
}
