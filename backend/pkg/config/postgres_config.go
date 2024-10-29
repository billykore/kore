package config

// postgres config.
type postgres struct {
	DSN string `envconfig:"POSTGRES_DSN"`
}
