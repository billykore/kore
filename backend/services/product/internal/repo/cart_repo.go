package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) List(ctx context.Context, username string, limit, startId int) ([]*model.Cart, error) {
	carts := make([]*model.Cart, 0)
	tx := r.db.WithContext(ctx).
		Preload("Product").
		Where("username = ?", username).
		Order("created_at DESC")
	if startId != 0 {
		tx = tx.Where("id > ?", startId)
	}
	tx = tx.Limit(limit).Find(&carts)
	return carts, tx.Error
}

func (r *CartRepository) Save(ctx context.Context, cart model.Cart) error {
	tx := r.db.WithContext(ctx).Save(&cart)
	return tx.Error
}

func (r *CartRepository) Update(ctx context.Context, id int, cart model.Cart) error {
	tx := r.db.WithContext(ctx).
		Model(&cart).
		Where("id = ?", id).
		Where("username = ?", cart.Username).
		Update("quantity", cart.Quantity)
	return tx.Error
}

func (r *CartRepository) Delete(ctx context.Context, id int, cart model.Cart) error {
	tx := r.db.WithContext(ctx).
		Where("id = ?", id).
		Where("username = ?", cart.Username).
		Delete(&cart)
	return tx.Error
}
