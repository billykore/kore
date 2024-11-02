package postgres

import (
	"context"

	"github.com/billykore/kore/backend/internal/domain/product"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) List(ctx context.Context, categoryId, limit, startId int) ([]*product.Product, error) {
	products := make([]*product.Product, 0)
	tx := r.db.WithContext(ctx).
		Preload("Discount").
		Preload("Category").
		Preload("Inventory")
	if categoryId > 0 {
		tx = tx.Where("category_id = ?", categoryId)
	}
	if startId > 0 {
		tx = tx.Where("id > ?", startId)
	}
	tx = tx.Order("id ASC").Limit(limit).Find(&products)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetById(ctx context.Context, id int) (*product.Product, error) {
	p := new(product.Product)
	tx := r.db.WithContext(ctx).
		Preload("Discount").
		Preload("Category").
		Preload("Inventory").
		Where("id = ?", id).
		First(p)
	return p, tx.Error
}

func (r *ProductRepository) CartList(ctx context.Context, username string, limit, startId int) ([]*product.Cart, error) {
	carts := make([]*product.Cart, 0)
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

func (r *ProductRepository) SaveCart(ctx context.Context, cart product.Cart) error {
	tx := r.db.WithContext(ctx).Save(&cart)
	return tx.Error
}

func (r *ProductRepository) UpdateCart(ctx context.Context, id int, cart product.Cart) error {
	tx := r.db.WithContext(ctx).
		Model(&cart).
		Where("id = ?", id).
		Where("username = ?", cart.Username).
		Update("quantity", cart.Quantity)
	return tx.Error
}

func (r *ProductRepository) DeleteCart(ctx context.Context, id int, cart product.Cart) error {
	tx := r.db.WithContext(ctx).
		Where("id = ?", id).
		Where("username = ?", cart.Username).
		Delete(&cart)
	return tx.Error
}

func (r *ProductRepository) CategoryList(ctx context.Context) ([]*product.Category, error) {
	categories := make([]*product.Category, 0)
	tx := r.db.WithContext(ctx).Find(&categories)
	return categories, tx.Error
}

func (r *ProductRepository) DiscountList(ctx context.Context, limit, startId int) ([]*product.Discount, error) {
	discounts := make([]*product.Discount, 0)
	tx := r.db.WithContext(ctx)
	if startId > 0 {
		tx = tx.Where("id > ?", startId)
	}
	tx = tx.Order("id ASC").Limit(limit).Find(&discounts)
	return discounts, tx.Error
}
