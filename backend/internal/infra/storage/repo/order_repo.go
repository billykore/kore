package repo

import (
	"context"

	"github.com/billykore/kore/backend/internal/domain/order"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) GetById(ctx context.Context, id uint) (*order.Order, error) {
	q := "SELECT id, username, payment_method, cart_ids, status FROM orders WHERE id = $1"
	tx := r.db.WithContext(ctx).Begin()
	return getOrder(tx, q, id)
}

func (r *OrderRepo) GetByIdAndStatus(ctx context.Context, id uint, status order.Status) (*order.Order, error) {
	q := "SELECT id, username, payment_method, cart_ids, status FROM orders WHERE id = $1 AND status = $2"
	tx := r.db.WithContext(ctx).Begin()
	return getOrder(tx, q, id, status)
}

func (r *OrderRepo) GetByShippingId(ctx context.Context, shippingId uint) (*order.Order, error) {
	q := "SELECT id, username, payment_method, cart_ids, status FROM orders WHERE shipping_id = $1"
	tx := r.db.WithContext(ctx).Begin()
	return getOrder(tx, q, shippingId)
}

func getOrder(tx *gorm.DB, query string, args ...any) (*order.Order, error) {
	o := new(order.Order)
	row := tx.Raw(query, args...).Row()
	err := row.Scan(&o.ID, &o.Username, &o.PaymentMethod, &o.CartIds, &o.Status)
	if err != nil {
		return nil, err
	}
	res := tx.Preload("Product").
		Where("id IN ?", o.IntCartIds()).
		Find(&o.Carts)
	if res.Error != nil {
		return nil, tx.Error
	}
	return o, nil
}

func (r *OrderRepo) Save(ctx context.Context, order order.Order) error {
	q := "INSERT INTO orders (username, payment_method, cart_ids, status) VALUES ($1, $2, $3, $4)"
	tx := r.db.WithContext(ctx).Begin()
	tx = tx.Exec(q, order.Username, order.PaymentMethod, order.CartIds, order.Status)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *OrderRepo) UpdateStatus(ctx context.Context, id uint, newStatus order.Status, currentStatus ...order.Status) error {
	q := "UPDATE orders SET status = ? WHERE id = ? AND status IN ?"
	tx := r.db.WithContext(ctx).Begin()
	err := updateOrder(tx, q, newStatus, id, currentStatus)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *OrderRepo) UpdateShipping(ctx context.Context, id uint, shippingId int) error {
	tx := r.db.WithContext(ctx).Begin()
	q1 := "UPDATE orders SET status = ? WHERE id = ?"
	err := updateOrder(tx, q1, order.StatusWaitingForShipment, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	q2 := "UPDATE orders SET shipping_id = ? WHERE id = ?"
	err = updateOrder(tx, q2, shippingId, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func updateOrder(tx *gorm.DB, query string, args ...any) error {
	res := tx.Exec(query, args...)
	return res.Error
}
