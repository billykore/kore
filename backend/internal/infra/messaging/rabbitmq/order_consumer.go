package rabbitmq

import (
	"context"

	"github.com/billykore/kore/backend/internal/domain/order"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/logger"
)

type OrderConsumer struct {
	cfg        *config.Config
	log        *logger.Logger
	svc        *order.Service
	connection *Connection
}

func NewOrderConsumer(cfg *config.Config, log *logger.Logger, connection *Connection, svc *order.Service) *OrderConsumer {
	return &OrderConsumer{
		cfg:        cfg,
		log:        log,
		svc:        svc,
		connection: connection,
	}
}

func (c *OrderConsumer) ListenOrderStatusChanges(ctx context.Context) error {
	msgs, err := c.connection.channel.ConsumeWithContext(ctx,
		c.cfg.Rabbit.QueueName, // queue
		"",                     // consumer
		true,                   // auto-ack
		false,                  // exclusive
		false,                  // no-local
		false,                  // no-wait
		nil,                    // args
	)
	if err != nil {
		c.log.Usecase("ListenOrderStatusChanges").Errorf("Consume error: %v", err)
		return err
	}

	go func() {
		for d := range msgs {
			c.log.Usecase("ListenOrderStatusChanges").
				Infof("Received a message: %s", d.Body)
			// Transform message and call the application service
			err := c.svc.ListenOrderStatusChanges(ctx, d.Body) // Example method in application service
			if err != nil {
				c.log.Usecase("ListenOrderStatusChanges").
					Infof("Failed to process user event: %v", err)
			}
		}
	}()

	return nil
}
