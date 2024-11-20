package internal

type Redis struct {
	Address  string `envconfig:"REDIS_ADDRESS"`
	Password string `envconfig:"REDIS_PASSWORD"`
	DB       int    `envconfig:"REDIS_DB"`
}
