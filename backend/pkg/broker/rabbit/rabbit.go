package rabbit

import (
	"context"

	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type HandlerFunc func(context.Context, amqp.Delivery) error

type Rabbit struct {
	log   *log.Logger
	conn  *amqp.Connection
	queue string
}

func New(cfg *config.Config, log *log.Logger, queue string) *Rabbit {
	conn, err := amqp.Dial(cfg.Rabbit.URL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	return &Rabbit{log: log, conn: conn, queue: queue}
}

func (r *Rabbit) Publish(ctx context.Context, body []byte) error {
	ch, err := r.conn.Channel()
	if err != nil {
		r.log.Fatalf("Failed to create channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(r.queue, true, false, false, false, nil)
	if err != nil {
		r.log.Fatalf("Failed to declare a queue: %v", err)
	}

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	return err
}

func (r *Rabbit) Consume(ctx context.Context, handler HandlerFunc) {
	ch, err := r.conn.Channel()
	if err != nil {
		r.log.Fatalf("Failed to create channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(r.queue, true, false, false, false, nil)
	if err != nil {
		r.log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.ConsumeWithContext(
		ctx,
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		r.log.Fatalf("Failed to create channel: %v", err)
	}

	forever := make(chan struct{})
	go func() {
		for msg := range msgs {
			err = handler(ctx, msg)
			if err != nil {
				r.log.Errorf("Failed to handle message: %v", err)
			}
		}
	}()

	<-forever
}

func (r *Rabbit) Close() {
	err := r.conn.Close()
	if err != nil {
		r.log.Errorf("Failed to close connection: %v", err)
	}
}
