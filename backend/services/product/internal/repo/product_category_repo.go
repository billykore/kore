package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"gorm.io/gorm"
)

type ProductCategoryRepo struct {
	db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) *ProductCategoryRepo {
	return &ProductCategoryRepo{db: db}
}

func (p *ProductCategoryRepo) List(ctx context.Context) ([]*model.ProductCategory, error) {
	categories := make([]*model.ProductCategory, 0)
	tx := p.db.WithContext(ctx).Find(&categories)
	return categories, tx.Error
}
