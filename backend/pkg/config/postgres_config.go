package config

// Postgres config.
type Postgres struct {
	DSN string `envconfig:"POSTGRES_DSN"`
}
