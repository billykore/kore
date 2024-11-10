package consumer

import (
	"context"

	"github.com/billykore/kore/backend/domain/order"
	"github.com/billykore/kore/backend/pkg/broker/rabbitmq"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/logger"
)

type OrderConsumer struct {
	cfg  *config.Config
	log  *logger.Logger
	svc  *order.Service
	conn *rabbitmq.Connection
}

func NewOrderConsumer(cfg *config.Config, log *logger.Logger, svc *order.Service, conn *rabbitmq.Connection) *OrderConsumer {
	return &OrderConsumer{
		cfg:  cfg,
		log:  log,
		svc:  svc,
		conn: conn,
	}
}

func (c *OrderConsumer) ListenOrderStatusChanges(ctx context.Context) error {
	msgs, err := c.conn.Channel.ConsumeWithContext(ctx,
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
		for msg := range msgs {
			c.log.Usecase("ListenOrderStatusChanges").Infof("Received a message: %s", msg.Body)

			payload := new(rabbitmq.MessagePayload[order.StatusChangeData])
			err := payload.UnmarshalBinary(msg.Body)
			if err != nil {
				c.log.Usecase("ListenOrderStatusChanges").
					Infof("Failed to process user event: %v", err)
			}

			err = c.svc.ConsumeOrderStatusChanges(ctx, payload.Data)
		}
	}()

	return nil
}
