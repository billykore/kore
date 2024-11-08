package http

import (
	handler2 "github.com/billykore/kore/backend/infra/http/handler"
	"github.com/billykore/kore/backend/infra/http/server"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	handler2.NewOrderHandler,
	handler2.NewOtpHandler,
	handler2.NewProductHandler,
	handler2.NewShippingHandler,
	handler2.NewUserHandler,
	server.NewRouter,
	server.New,
)
