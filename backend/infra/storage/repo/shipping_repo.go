package repo

import (
	"context"

	"github.com/billykore/kore/backend/domain/shipping"
	"gorm.io/gorm"
)

type ShippingRepo struct {
	db *gorm.DB
}

func NewShippingRepo(db *gorm.DB) *ShippingRepo {
	return &ShippingRepo{
		db: db,
	}
}

func (r *ShippingRepo) GetById(ctx context.Context, id uint) (*shipping.Shipping, error) {
	s := new(shipping.Shipping)
	tx := r.db.WithContext(ctx).Where("id = ?", id).First(s)
	return s, tx.Error
}

func (r *ShippingRepo) Save(ctx context.Context, shipping shipping.Shipping) (uint, error) {
	tx := r.db.WithContext(ctx).Save(&shipping)
	return shipping.ID, tx.Error
}

func (r *ShippingRepo) UpdateStatus(ctx context.Context, id uint, newStatus, currentStatus shipping.Status) error {
	tx := r.db.WithContext(ctx).
		Model(&shipping.Shipping{}).
		Where("id = ?", id).
		Where("status = ?", currentStatus).
		UpdateColumn("status", newStatus)
	return tx.Error
}
