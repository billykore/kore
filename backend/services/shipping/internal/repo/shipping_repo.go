package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
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

func (r *ShippingRepository) GetById(ctx context.Context, id uint) (*model.Shipping, error) {
	shipping := new(model.Shipping)
	tx := r.db.WithContext(ctx).Where("id = ?", id).First(shipping)
	return shipping, tx.Error
}

func (r *ShippingRepository) Save(ctx context.Context, shipping model.Shipping) (uint, error) {
	tx := r.db.WithContext(ctx).Save(&shipping)
	return shipping.ID, tx.Error
}

func (r *ShippingRepository) UpdateStatus(ctx context.Context, id uint, newStatus, currentStatus model.ShippingStatus) error {
	tx := r.db.WithContext(ctx).
		Model(&model.Shipping{}).
		Where("id = ?", id).
		Where("status = ?", currentStatus).
		UpdateColumn("status", newStatus)
	return tx.Error
}
