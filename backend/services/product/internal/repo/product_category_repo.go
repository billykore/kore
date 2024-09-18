package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"gorm.io/gorm"
)

type ProductCategoryRepository struct {
	db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) *ProductCategoryRepository {
	return &ProductCategoryRepository{db: db}
}

func (r *ProductCategoryRepository) List(ctx context.Context) ([]*model.ProductCategory, error) {
	categories := make([]*model.ProductCategory, 0)
	tx := r.db.WithContext(ctx).Find(&categories)
	return categories, tx.Error
}
