package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Connection struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewConnection establishes a new RabbitMQ connection.
func NewConnection(dsn string) (*Connection, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Connection{
		conn:    conn,
		channel: channel,
	}, nil
}

// Close closes the RabbitMQ connection and channel.
func (c *Connection) Close() {
	err := c.channel.Close()
	if err != nil {
		return
	}
	err = c.conn.Close()
	if err != nil {
		return
	}
}
