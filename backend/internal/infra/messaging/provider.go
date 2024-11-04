package messaging

import (
	"github.com/billykore/kore/backend/internal/domain/shipping"
	"github.com/billykore/kore/backend/internal/infra/messaging/consumer"
	"github.com/billykore/kore/backend/internal/infra/messaging/producer"
	"github.com/billykore/kore/backend/internal/infra/messaging/rabbitmq"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	producer.NewShippingProducer, wire.Bind(new(shipping.Messaging), new(*producer.ShippingProducer)),
	consumer.NewOrderConsumer,
	rabbitmq.NewConnection,
	NewConsumer,
)
