package repo

import (
	"context"
	"testing"

	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/db"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetOrderById(t *testing.T) {
	err := godotenv.Load("../../../.env")
	assert.NoError(t, err)

	cfg := config.Get()
	pg := db.NewPostgres(cfg)
	repo := &orderRepo{db: pg}

	order, err := repo.GetById(context.Background(), 2)
	assert.NoError(t, err)
	assert.NotEmpty(t, order)

	t.Log(order)
}
