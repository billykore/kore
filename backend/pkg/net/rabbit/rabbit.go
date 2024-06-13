package rabbit

import (
	"context"

	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Rabbit is the message broker to communicate data asynchronously.
type Rabbit struct {
	log   *log.Logger
	conn  *amqp.Connection
	queue string
}

// New initiate new Rabbit.
func New(cfg *config.Config, log *log.Logger) *Rabbit {
	conn, err := amqp.Dial(cfg.Rabbit.URL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	return &Rabbit{log: log, conn: conn, queue: cfg.Rabbit.QueueName}
}

// Publish will publish the message to Rabbit.
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
	if err != nil {
		r.log.Fatalf("Failed to publish a message: %v", err)
		return err
	}

	r.log.Infof("Published a message: %s", string(body))
	return nil
}

// HandlerFunc is function to handle operations when consume a message from Rabbit.
type HandlerFunc func(context.Context, amqp.Delivery) error

// Consume consumes the message from message broker and the message is handled by HandlerFunc.
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
			r.log.Infof("Consumed message: %s", string(msg.Body))
		}
	}()

	<-forever
}

// Close the Rabbit connection.
func (r *Rabbit) Close() {
	err := r.conn.Close()
	if err != nil {
		r.log.Errorf("Failed to close connection: %v", err)
	}
}
