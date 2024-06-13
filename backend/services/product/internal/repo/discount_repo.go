package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"gorm.io/gorm"
)

type DiscountRepo struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) *DiscountRepo {
	return &DiscountRepo{
		db: db,
	}
}

func (r *DiscountRepo) List(ctx context.Context, limit, startId int) ([]*model.Discount, error) {
	discounts := make([]*model.Discount, 0)
	tx := r.db.WithContext(ctx)
	if startId > 0 {
		tx = tx.Where("id > ?", startId)
	}
	tx = tx.Order("id ASC").Limit(limit).Find(&discounts)
	return discounts, tx.Error
}
