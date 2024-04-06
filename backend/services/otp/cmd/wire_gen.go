// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/db"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/mail"
	"github.com/billykore/kore/backend/services/otp/internal/handler"
	"github.com/billykore/kore/backend/services/otp/internal/repo"
	"github.com/billykore/kore/backend/services/otp/internal/server"
	"github.com/billykore/kore/backend/services/otp/internal/usecase"
	"github.com/labstack/echo/v4"
)

import (
	_ "github.com/joho/godotenv/autoload"
)

// Injectors from wire.go:

func otpApp(cfg *config.Config) *app {
	logger := log.NewLogger()
	echoEcho := echo.New()
	gormDB := db.NewPostgres(cfg)
	otpRepo := repo.NewOtpRepository(gormDB)
	sender := mail.NewSender(cfg)
	otpUsecase := usecase.NewOtpUsecase(logger, otpRepo, sender)
	otpHandler := handler.NewOtpHandler(otpUsecase)
	router := server.NewRouter(cfg, logger, echoEcho, otpHandler)
	httpServer := server.NewHTTPServer(router)
	mainApp := newApp(httpServer)
	return mainApp
}
