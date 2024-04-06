package db

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgres returns postgres db connection instance.
func NewPostgres(cfg *config.Config) *gorm.DB {
	dsn := cfg.Postgres.DSN
	logger := log.NewLogger()
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		logger.Usecase("NewPostgres").Fatalf("failed to connect database: %v", err)
	}
	return db
}
