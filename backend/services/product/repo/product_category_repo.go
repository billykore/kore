package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
	"gorm.io/gorm"
)

type productCategoryRepo struct {
	db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) repo.ProductCategoryRepository {
	return &productCategoryRepo{db: db}
}

func (p *productCategoryRepo) List(ctx context.Context) ([]*model.ProductCategory, error) {
	categories := make([]*model.ProductCategory, 0)
	tx := p.db.WithContext(ctx).Find(&categories)
	return categories, tx.Error
}
