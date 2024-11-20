// Package config contains all the services configuration values.
// The configuration is from .env file.
package config

import (
	"sync"

	"github.com/billykore/kore/backend/pkg/config/internal"
	"github.com/kelseyhightower/envconfig"
)

// Config contains app configurations use by services.
type Config struct {
	HTTPPort  string `envconfig:"HTTP_PORT" default:"8000"`
	Postgres  internal.Postgres
	Firestore internal.Firestore
	Token     internal.Token
	Rabbit    internal.Rabbit
	Email     internal.Email
	Redis     internal.Redis
}

var (
	_cfg  *Config
	_once sync.Once
)

// Get return the singleton instance of Config.
func Get() *Config {
	_once.Do(func() {
		_cfg = new(Config)
		envconfig.MustProcess("", _cfg)
	})
	return _cfg
}
