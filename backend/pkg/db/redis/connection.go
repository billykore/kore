package redis

import (
	"context"
	"time"

	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/redis/go-redis/v9"
)

func New(cfg *config.Config) *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:            cfg.Redis.Address,
		Password:        cfg.Redis.Password,
		DB:              cfg.Redis.DB,
		PoolSize:        5,
		MinIdleConns:    3,
		MaxIdleConns:    5,
		MaxActiveConns:  10,
		ReadTimeout:     3 * time.Second,
		WriteTimeout:    3 * time.Second,
		ConnMaxIdleTime: 30 * time.Minute,
	})
	cmd := r.Ping(context.Background())
	if err := cmd.Err(); err != nil {
		logger.New().Fatalf("could not connect to redis: %v", err)
		return nil
	}
	return r
}
