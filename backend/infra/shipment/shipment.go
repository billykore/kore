package shipment

import (
	"github.com/billykore/kore/backend/domain/order"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewJNE, wire.Bind(new(order.ShippingService), new(*JNE)),
)
