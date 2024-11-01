package rabbitmq

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ShippingProducer struct {
	log        *logger.Logger
	connection *Connection
	exchange   string
}

func NewShippingProducer(connection *Connection, exchange string) *ShippingProducer {
	return &ShippingProducer{
		connection: connection,
		exchange:   exchange,
	}
}

type UpdateShippingRabbitData struct {
	ShippingId uint   `json:"shippingId"`
	Status     string `json:"status"`
}

func (p *ShippingProducer) PublishShippingUpdateStatus(payload entity.MessagePayload[*UpdateShippingRabbitData]) error {
	body, err := payload.MarshalBinary()
	if err != nil {
		return err
	}
	// Publish message to the RabbitMQ exchange
	err = p.connection.channel.Publish(
		p.exchange, // exchange
		"",         // routing key
		false,      // mandatory
		false,      // immediate
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
