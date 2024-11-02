package rabbitmq

import (
	"context"

	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ShippingProducer struct {
	cfg        *config.Config
	log        *logger.Logger
	connection *Connection
}

func NewShippingProducer(cfg *config.Config, connection *Connection) *ShippingProducer {
	return &ShippingProducer{
		cfg:        cfg,
		connection: connection,
	}
}

type UpdateShippingRabbitData struct {
	ShippingId uint   `json:"shippingId"`
	Status     string `json:"status"`
}

func (p *ShippingProducer) PublishShippingUpdateStatus(ctx context.Context, payload *entity.MessagePayload[*UpdateShippingRabbitData]) error {
	body, err := payload.MarshalBinary()
	if err != nil {
		return err
	}
	// Publish message to the RabbitMQ exchange
	err = p.connection.channel.PublishWithContext(ctx,
		p.cfg.Rabbit.QueueName, // exchange
		"",                     // routing key
		false,                  // mandatory
		false,                  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		p.log.Errorf("Failed to publish message: %v", err)
		return err
	}
	p.log.Infof("Published shipping status update for shippingId: %d", payload.Data.ShippingId)
	return nil
}
