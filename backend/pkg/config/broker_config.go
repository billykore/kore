package config

type Rabbit struct {
	URL string `envconfig:"RABBIT_URL"`
}
