package payment

import (
	"github.com/billykore/kore/backend/internal/domain/order"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewGoPay, wire.Bind(new(order.PaymentService), new(*GoPay)),
)
