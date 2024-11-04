package repository

import (
	"context"

	"github.com/billykore/kore/backend/internal/domain/shipping"
	"gorm.io/gorm"
)

type ShippingRepository struct {
	db *gorm.DB
}

func NewShippingRepository(db *gorm.DB) *ShippingRepository {
	return &ShippingRepository{
		db: db,
	}
}

func (r *ShippingRepository) GetById(ctx context.Context, id uint) (*shipping.Shipping, error) {
	s := new(shipping.Shipping)
	tx := r.db.WithContext(ctx).Where("id = ?", id).First(s)
	return s, tx.Error
}

func (r *ShippingRepository) Save(ctx context.Context, shipping shipping.Shipping) (uint, error) {
	tx := r.db.WithContext(ctx).Save(&shipping)
	return shipping.ID, tx.Error
}

func (r *ShippingRepository) UpdateStatus(ctx context.Context, id uint, newStatus, currentStatus shipping.Status) error {
	tx := r.db.WithContext(ctx).
		Model(&shipping.Shipping{}).
		Where("id = ?", id).
		Where("status = ?", currentStatus).
		UpdateColumn("status", newStatus)
	return tx.Error
}
