package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

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

func Get() *Config {
	_once.Do(func() {
		_cfg = new(Config)
		envconfig.MustProcess("", _cfg)
	})
	return _cfg
}
