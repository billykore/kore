package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port      string `envconfig:"PORT"`
	Firestore Firestore
}

var (
	cfg  *Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		cfg = new(Config)
		envconfig.MustProcess("", cfg)
	})
	return cfg
}
