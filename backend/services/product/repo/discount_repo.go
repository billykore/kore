package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
	"gorm.io/gorm"
)

type discountRepo struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) repo.DiscountRepository {
	return &discountRepo{
		db: db,
	}
}

func (r *discountRepo) List(ctx context.Context, limit, startId int) ([]*model.Discount, error) {
	discounts := make([]*model.Discount, 0)
	tx := r.db.WithContext(ctx)
	if startId > 0 {
		tx = tx.Where("id > ?", startId)
	}
	tx = tx.Order("id ASC").Limit(limit).Find(&discounts)
	return discounts, tx.Error
}
