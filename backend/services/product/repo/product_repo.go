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

func (r *productRepo) List(ctx context.Context, categoryId, limit, startId int) ([]*model.Product, error) {
	products := make([]*model.Product, 0)
	tx := r.db.WithContext(ctx).
		Preload("Discount").
		Preload("ProductCategory").
		Preload("ProductInventory")
	if categoryId > 0 {
		tx = tx.Where("category_id = ?", categoryId)
	}
	if startId > 0 {
		tx = tx.Where("id > ?", startId)
	}
	tx = tx.Order("id ASC").Limit(limit).Find(&products)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepo) GetById(ctx context.Context, id int) (*model.Product, error) {
	product := new(model.Product)
	tx := r.db.WithContext(ctx).
		Preload("Discount").
		Preload("ProductCategory").
		Preload("ProductInventory").
		Where("id = ?", id).
		First(product)
	return product, tx.Error
}
