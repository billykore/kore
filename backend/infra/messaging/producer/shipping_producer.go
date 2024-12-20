package producer

import (
	"context"

	"github.com/billykore/kore/backend/domain/shipping"
	"github.com/billykore/kore/backend/pkg/broker/rabbitmq"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ShippingProducer struct {
	cfg  *config.Config
	log  *logger.Logger
	conn *rabbitmq.Connection
}

func NewShippingProducer(cfg *config.Config, log *logger.Logger, conn *rabbitmq.Connection) *ShippingProducer {
	return &ShippingProducer{
		cfg:  cfg,
		log:  log,
		conn: conn,
	}
}

func (p *ShippingProducer) PublishStatusChange(ctx context.Context, data shipping.StatusChangeData) error {
	payload := rabbitmq.MessagePayload[shipping.StatusChangeData]{
		Origin: "shipping-service",
		Data:   data,
	}
	body, err := payload.MarshalBinary()
	if err != nil {
		return err
	}

	err = p.conn.Channel.ExchangeDeclare(
		p.cfg.Rabbit.QueueName, // name
		"fanout",               // type
		true,                   // durable
		false,                  // auto-deleted
		false,                  // internal
		false,                  // no-wait
		nil,                    // arguments
	)

	// Publish message to the RabbitMQ exchange
	err = p.conn.Channel.PublishWithContext(ctx,
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
