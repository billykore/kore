package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config is app configurations use by services.
type Config struct {
	HTTPPort  string `envconfig:"HTTP_PORT" default:"8000"`
	Postgres  Postgres
	Firestore Firestore
	Token     Token
	Rabbit    Rabbit
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
