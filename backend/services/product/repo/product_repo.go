package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repo.ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) List(ctx context.Context) ([]model.Product, error) {
	products := make([]model.Product, 0)
	res := r.db.WithContext(ctx).
		Preload("Discount").
		Preload("ProductCategory").
		Preload("ProductInventory").
		Find(&products)
	if err := res.Error; err != nil {
		return nil, err
	}
	return products, nil
}
