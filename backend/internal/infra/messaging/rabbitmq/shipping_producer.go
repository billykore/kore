package rabbitmq

import (
	"context"

	"github.com/billykore/kore/backend/internal/domain/shipping"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ShippingProducer struct {
	cfg  *config.Config
	log  *logger.Logger
	conn *Connection
	svc  *shipping.Service
}

func NewShippingProducer(cfg *config.Config, conn *Connection) *ShippingProducer {
	return &ShippingProducer{
		cfg:  cfg,
		conn: conn,
	}
}

func (p *ShippingProducer) ProduceStatusChange(ctx context.Context, data shipping.StatusChangeData) error {
	payload := entity.MessagePayload[shipping.StatusChangeData]{
		Origin: "shipping-service",
		Data:   data,
	}
	body, err := payload.MarshalBinary()
	if err != nil {
		return err
	}
	// Publish message to the RabbitMQ exchange
	err = p.conn.channel.PublishWithContext(ctx,
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
