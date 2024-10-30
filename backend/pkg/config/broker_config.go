package config

// rabbit configuration.
type rabbit struct {
	URL       string `envconfig:"RABBIT_URL"`
	QueueName string `envconfig:"RABBIT_QUEUE_NAME"`
}
