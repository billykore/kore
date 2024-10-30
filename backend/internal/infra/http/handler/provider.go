package handler

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewOrderHandler,
	NewOtpHandler,
	NewProductHandler,
	NewShippingHandler,
	NewUserHandler,
)
