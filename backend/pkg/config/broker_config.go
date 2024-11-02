package config

// rabbit configuration.
type rabbit struct {
	DSN       string `envconfig:"RABBIT_DSN"`
	QueueName string `envconfig:"RABBIT_QUEUE_NAME"`
}
