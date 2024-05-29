package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) GetById(ctx context.Context, id uint) (*model.Order, error) {
	q := "SELECT id, user_id, payment_method, cart_ids, status FROM orders WHERE id = $1"
	tx := r.db.WithContext(ctx).Begin()
	return getOrder(tx, q, id)
}

func (r *OrderRepo) GetByIdAndStatus(ctx context.Context, id uint, status model.OrderStatus) (*model.Order, error) {
	q := "SELECT id, user_id, payment_method, cart_ids, status FROM orders WHERE id = $1 AND status = $2"
	tx := r.db.WithContext(ctx).Begin()
	return getOrder(tx, q, id, status)
}

func (r *OrderRepo) GetByShippingId(ctx context.Context, shippingId uint) (*model.Order, error) {
	q := "SELECT id, user_id, payment_method, cart_ids, status FROM orders WHERE shipping_id = $1"
	tx := r.db.WithContext(ctx).Begin()
	return getOrder(tx, q, shippingId)
}

func getOrder(tx *gorm.DB, query string, args ...any) (*model.Order, error) {
	order := new(model.Order)
	row := tx.Raw(query, args...).Row()
	err := row.Scan(&order.ID, &order.UserId, &order.PaymentMethod, &order.CartIds, &order.Status)
	if err != nil {
		return nil, err
	}
	res := tx.Preload("Product").
		Where("id IN ?", order.IntCartIds()).
		Find(&order.Carts)
	if res.Error != nil {
		return nil, tx.Error
	}
	return order, nil
}

func (r *OrderRepo) Save(ctx context.Context, order model.Order) error {
	tx := r.db.WithContext(ctx).Save(&order)
	return tx.Error
}

func (r *OrderRepo) UpdateStatus(ctx context.Context, id uint, newStatus model.OrderStatus, currentStatus ...model.OrderStatus) error {
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
	err := updateOrder(tx, q1, model.OrderStatusWaitingForShipment, id)
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
