package messaging

import (
	"github.com/billykore/kore/backend/internal/infra/messaging/rabbitmq"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	rabbitmq.NewShippingProducer,
	rabbitmq.NewOrderConsumer,
	rabbitmq.NewConnection,
	NewConsumers,
)
