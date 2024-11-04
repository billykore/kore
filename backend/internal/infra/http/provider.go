package http

import (
	"github.com/billykore/kore/backend/internal/infra/http/handler"
	"github.com/billykore/kore/backend/internal/infra/http/server"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	handler.NewOrderHandler,
	handler.NewOtpHandler,
	handler.NewProductHandler,
	handler.NewShippingHandler,
	handler.NewUserHandler,
	server.NewRouter,
	server.New,
)
