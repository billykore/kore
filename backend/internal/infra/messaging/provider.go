package messaging

import (
	"github.com/billykore/kore/backend/internal/domain/shipping"
	"github.com/billykore/kore/backend/internal/infra/messaging/rabbitmq"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	rabbitmq.NewShippingProducer, wire.Bind(new(shipping.Messaging), new(*rabbitmq.ShippingProducer)),
	rabbitmq.NewOrderConsumer,
	rabbitmq.NewConnection,
	NewConsumers,
)
