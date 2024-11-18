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
	err := c.conn.Channel.ExchangeDeclare(
		c.cfg.Rabbit.QueueName, // name
		"fanout",               // type
		true,                   // durable
		false,                  // auto-deleted
		false,                  // internal
		false,                  // no-wait
		nil,                    // arguments
	)
	if err != nil {
		c.log.Usecase("ListenOrderStatusChanges").
			Errorf("ExchangeDeclare error: %v", err)
		return err
	}

	queue, err := c.conn.Channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		c.log.Usecase("ListenOrderStatusChanges").Errorf("QueueDeclare error: %v", err)
		return err
	}

	err = c.conn.Channel.QueueBind(
		queue.Name,             // queue name
		"",                     // routing key
		c.cfg.Rabbit.QueueName, // exchange
		false,
		nil,
	)
	if err != nil {
		c.log.Usecase("ListenOrderStatusChanges").
			Errorf("QueueBind error: %v", err)
	}

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
		c.log.Usecase("ListenOrderStatusChanges").
			Errorf("ConsumeWithContext error: %v", err)
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
				return
			}

			err = c.svc.ConsumeOrderStatusChanges(ctx, payload.Data)
			if err != nil {
				c.log.Usecase("ListenOrderStatusChanges").
					Errorf("Failed to process user event: %v", err)
				return
			}
		}
	}()

	return nil
}
