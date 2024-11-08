package messaging

import (
	"github.com/billykore/kore/backend/domain/shipping"
	"github.com/billykore/kore/backend/infra/messaging/consumer"
	"github.com/billykore/kore/backend/infra/messaging/producer"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	producer.NewShippingProducer, wire.Bind(new(shipping.Messaging), new(*producer.ShippingProducer)),
	consumer.NewOrderConsumer,
	NewConsumer,
)
