package repo

import (
	"context"
	"fmt"
	"testing"

	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/db"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestProductRepoList(t *testing.T) {
	err := godotenv.Load("../../../.env")
	assert.NoError(t, err)

	cfg := config.Get()
	pg := db.NewPostgres(cfg)

	repo := productRepo{db: pg}
	products, err := repo.List(context.Background())
	assert.NoError(t, err)

	for _, product := range products {
		fmt.Printf("%+v\n", product.ProductInventory)
		fmt.Printf("%+v\n", product.ProductCategory)
		fmt.Printf("%+v\n", product.Discount)
	}
}
