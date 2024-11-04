package postgres

import (
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New returns new postgres db connection.
func New(cfg *config.Config) *gorm.DB {
	dsn := cfg.Postgres.DSN
	log := logger.New()
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Usecase("New").Fatalf("failed to connect database: %v", err)
	}
	return db
}
