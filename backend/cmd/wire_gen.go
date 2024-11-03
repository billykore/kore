// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/billykore/kore/backend/internal/domain/order"
	"github.com/billykore/kore/backend/internal/domain/otp"
	"github.com/billykore/kore/backend/internal/domain/product"
	"github.com/billykore/kore/backend/internal/domain/shipping"
	"github.com/billykore/kore/backend/internal/domain/user"
	"github.com/billykore/kore/backend/internal/infra/email"
	"github.com/billykore/kore/backend/internal/infra/http"
	"github.com/billykore/kore/backend/internal/infra/http/handler"
	"github.com/billykore/kore/backend/internal/infra/messaging"
	"github.com/billykore/kore/backend/internal/infra/messaging/rabbitmq"
	"github.com/billykore/kore/backend/internal/infra/persistence/postgres"
	postgres2 "github.com/billykore/kore/backend/internal/infra/storage/postgres"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/validation"
	"github.com/labstack/echo/v4"
)

import (
	_ "github.com/joho/godotenv/autoload"
)

// Injectors from wire.go:

func initApp(cfg *config.Config) *app {
	loggerLogger := logger.New()
	echoEcho := echo.New()
	db := postgres.New(cfg)
	userRepository := postgres2.NewUserRepository(db)
	service := user.NewService(loggerLogger, userRepository)
	userHandler := handler.NewUserHandler(service)
	orderRepository := postgres2.NewOrderRepository(db)
	orderService := order.NewService(loggerLogger, orderRepository)
	connection := rabbitmq.NewConnection(cfg)
	orderHandler := handler.NewOrderHandler(orderService, connection)
	otpRepository := postgres2.NewOtpRepository(db)
	client := email.NewClient(cfg)
	otpEmail := email.NewOTPEmail(loggerLogger, client)
	otpService := otp.NewService(loggerLogger, otpRepository, otpEmail)
	validator := validation.New()
	otpHandler := handler.NewOtpHandler(otpService, validator)
	productRepository := postgres2.NewProductRepository(db)
	productService := product.NewService(loggerLogger, productRepository)
	productHandler := handler.NewProductHandler(productService)
	shippingRepository := postgres2.NewShippingRepository(db)
	shippingService := shipping.NewService(loggerLogger, shippingRepository)
	shippingProducer := rabbitmq.NewShippingProducer(cfg, connection)
	shippingHandler := handler.NewShippingHandler(shippingService, shippingProducer)
	router := http.NewRouter(cfg, loggerLogger, echoEcho, userHandler, orderHandler, otpHandler, productHandler, shippingHandler)
	server := http.NewServer(router)
	orderConsumer := rabbitmq.NewOrderConsumer(cfg, loggerLogger, connection, orderService)
	consumer := messaging.NewConsumers(cfg, loggerLogger, orderConsumer)
	mainApp := newApp(server, consumer)
	return mainApp
}
