package shipping

import (
	"github.com/billykore/kore/backend/internal/domain/order"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewJNE, wire.Bind(new(order.ShippingService), new(*JNE)),
)
