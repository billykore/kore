package payment

import (
	"github.com/billykore/kore/backend/domain/order"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewGoPay, wire.Bind(new(order.PaymentService), new(*GoPay)),
)
