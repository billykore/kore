package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
	"gorm.io/gorm"
)

type shippingRepo struct {
	db *gorm.DB
}

func NewShippingRepository(db *gorm.DB) repo.ShippingRepository {
	return &shippingRepo{
		db: db,
	}
}

func (r *shippingRepo) Save(ctx context.Context, shipping model.Shipping) (uint, error) {
	tx := r.db.WithContext(ctx).Save(&shipping)
	return shipping.ID, tx.Error
}
