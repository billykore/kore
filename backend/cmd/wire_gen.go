// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/billykore/kore/backend/domain/order"
	"github.com/billykore/kore/backend/domain/otp"
	"github.com/billykore/kore/backend/domain/product"
	"github.com/billykore/kore/backend/domain/shipping"
	"github.com/billykore/kore/backend/domain/user"
	"github.com/billykore/kore/backend/infra/email/mailer"
	"github.com/billykore/kore/backend/infra/http/handler"
	"github.com/billykore/kore/backend/infra/http/server"
	"github.com/billykore/kore/backend/infra/messaging"
	"github.com/billykore/kore/backend/infra/messaging/consumer"
	"github.com/billykore/kore/backend/infra/messaging/producer"
	"github.com/billykore/kore/backend/infra/payment"
	"github.com/billykore/kore/backend/infra/shipment"
	"github.com/billykore/kore/backend/infra/storage/repo"
	"github.com/billykore/kore/backend/pkg/broker/rabbitmq"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/db/postgres"
	"github.com/billykore/kore/backend/pkg/email/brevo"
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
	validator := validation.New()
	db := postgres.New(cfg)
	userRepo := repo.NewUserRepo(db)
	service := user.NewService(loggerLogger, userRepo)
	userHandler := handler.NewUserHandler(validator, service)
	orderRepo := repo.NewOrderRepo(db)
	goPay := payment.NewGoPay()
	jne := shipment.NewJNE()
	orderService := order.NewService(loggerLogger, orderRepo, goPay, jne)
	connection := rabbitmq.NewConnection(cfg)
	orderHandler := handler.NewOrderHandler(orderService, connection)
	otpRepo := repo.NewOtpRepo(db)
	client := brevo.NewClient(cfg)
	otpEmail := mailer.NewOTPEmail(loggerLogger, client)
	otpService := otp.NewService(loggerLogger, otpRepo, otpEmail)
	otpHandler := handler.NewOtpHandler(otpService, validator)
	productRepo := repo.NewProductRepo(db)
	productService := product.NewService(loggerLogger, productRepo)
	productHandler := handler.NewProductHandler(productService)
	shippingRepo := repo.NewShippingRepo(db)
	shippingProducer := producer.NewShippingProducer(cfg, loggerLogger, connection)
	shippingService := shipping.NewService(loggerLogger, shippingRepo, shippingProducer)
	shippingHandler := handler.NewShippingHandler(shippingService)
	router := server.NewRouter(cfg, loggerLogger, echoEcho, userHandler, orderHandler, otpHandler, productHandler, shippingHandler)
	serverServer := server.New(router)
	orderConsumer := consumer.NewOrderConsumer(cfg, loggerLogger, orderService, connection)
	messagingConsumer := messaging.NewConsumer(cfg, loggerLogger, orderConsumer)
	mainApp := newApp(serverServer, messagingConsumer)
	return mainApp
}
