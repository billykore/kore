package rabbit

import (
	"context"
	"testing"

	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

func TestRabbitPublish(t *testing.T) {
	err := godotenv.Load("../../../.env")
	assert.NoError(t, err)
	cfg := config.Get()
	logg := log.NewLogger()

	rabbit := New(cfg, logg)
	assert.NotNil(t, rabbit)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	payload := &Payload[*entity.UpdateShippingRabbitData]{
		Origin: "shipping-service",
		Data: &entity.UpdateShippingRabbitData{
			ShippingId: 999,
			Status:     "shipped",
		},
	}
	bytePayload, err := payload.MarshalBinary()

	err = rabbit.Publish(ctx, bytePayload)
	assert.NoError(t, err)
}

func TestRabbitConsume(t *testing.T) {
	err := godotenv.Load("../../../.env")
	assert.NoError(t, err)
	cfg := config.Get()
	logg := log.NewLogger()

	rabbit := New(cfg, logg)
	assert.NotNil(t, rabbit)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rabbit.Consume(ctx, func(ctx context.Context, delivery amqp.Delivery) error {
		t.Log(string(delivery.Body))
		return nil
	})
}
