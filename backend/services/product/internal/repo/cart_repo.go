package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"gorm.io/gorm"
)

type CartRepo struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepo {
	return &CartRepo{
		db: db,
	}
}

func (r *CartRepo) List(ctx context.Context, userId, limit, startId int) ([]*model.Cart, error) {
	carts := make([]*model.Cart, 0)
	tx := r.db.WithContext(ctx).
		Preload("Product").
		Where("user_id = ?", userId).
		Order("created_at DESC")
	if startId != 0 {
		tx = tx.Where("id > ?", startId)
	}
	tx = tx.Limit(limit).Find(&carts)
	return carts, tx.Error
}

func (r *CartRepo) Save(ctx context.Context, item model.Cart) error {
	tx := r.db.WithContext(ctx).Save(&item)
	return tx.Error
}

func (r *CartRepo) Update(ctx context.Context, id, quantity int) error {
	tx := r.db.WithContext(ctx).
		Model(&model.Cart{}).
		Where("id = ?", id).
		Update("quantity", quantity)
	return tx.Error
}

func (r *CartRepo) Delete(ctx context.Context, id int) error {
	tx := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Cart{})
	return tx.Error
}
