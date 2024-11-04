package internal

// Email config.
type Email struct {
	From string `envconfig:"EMAIL_FROM"`
	Host string `envconfig:"EMAIL_HOST"`
	Port int    `envconfig:"EMAIL_PORT"`
	Key  string `envconfig:"EMAIL_KEY"`
}
