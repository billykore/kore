package rabbitmq

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Connection struct {
	conn    *amqp.Connection
	Channel *amqp.Channel
}

// NewConnection establishes a new RabbitMQ connection.
func NewConnection(cfg *config.Config) *Connection {
	log := logger.New()

	conn, err := amqp.Dial(cfg.Rabbit.DSN)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		return nil
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a Channel: %s", err)
		return nil
	}
	return &Connection{
		conn:    conn,
		Channel: channel,
	}
}

// Close closes the RabbitMQ connection and channel.
func (c *Connection) Close() {
	err := c.Channel.Close()
	if err != nil {
		return
	}
	err = c.conn.Close()
	if err != nil {
		return
	}
}
