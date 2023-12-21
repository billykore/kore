package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env       string `envconfig:"ENV"`
	Port      string `envconfig:"PORT"`
	Firestore Firestore
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
