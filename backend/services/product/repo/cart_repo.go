package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
	"gorm.io/gorm"
)

type cartRepo struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) repo.CartRepository {
	return &cartRepo{
		db: db,
	}
}

func (r *cartRepo) List(ctx context.Context, userId, limit, startId int) ([]*model.Cart, error) {
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

func (r *cartRepo) Save(ctx context.Context, item model.Cart) error {
	tx := r.db.WithContext(ctx).Save(&item)
	return tx.Error
}

func (r *cartRepo) Update(ctx context.Context, id, quantity int) error {
	tx := r.db.WithContext(ctx).
		Model(&model.Cart{}).
		Where("id = ?", id).
		Update("quantity", quantity)
	return tx.Error
}

func (r *cartRepo) Delete(ctx context.Context, id int) error {
	tx := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Cart{})
	return tx.Error
}
