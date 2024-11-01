package rabbitmq

import (
	"context"

	"github.com/billykore/kore/backend/internal/app/order"
	"github.com/billykore/kore/backend/pkg/logger"
)

type OrderConsumer struct {
	log        *logger.Logger
	svc        *order.Service
	connection *Connection
	queue      string
}

func NewOrderConsumer(log *logger.Logger, connection *Connection, svc *order.Service, queue string) *OrderConsumer {
	return &OrderConsumer{
		log:        log,
		svc:        svc,
		connection: connection,
		queue:      queue,
	}
}

func (c *OrderConsumer) ConsumeOrderStatusChanges(ctx context.Context) error {
	msgs, err := c.connection.channel.ConsumeWithContext(ctx,
		c.queue, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		c.log.Usecase("ConsumeOrderStatusChanges").Errorf("Consume error: %v", err)
		return err
	}

	go func() {
		for d := range msgs {
			c.log.Usecase("ConsumeOrderStatusChanges").
				Infof("Received a message: %s", d.Body)
			// Transform message and call the application service
			err := c.svc.ListenOrderStatusChanges(ctx, d.Body) // Example method in application service
			if err != nil {
				c.log.Usecase("ConsumeOrderStatusChanges").
					Infof("Failed to process user event: %v", err)
			}
		}
	}()

	return nil
}
