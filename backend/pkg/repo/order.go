package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

type OrderRepository interface {
	GetById(ctx context.Context, id uint) (*model.Order, error)
	GetByIdAndStatus(ctx context.Context, id uint, status model.OrderStatus) (*model.Order, error)
	GetByShippingId(ctx context.Context, shippingId uint) (*model.Order, error)
	Save(ctx context.Context, order model.Order) error
	UpdateStatus(ctx context.Context, id uint, newStatus model.OrderStatus, currentStatus ...model.OrderStatus) error
	UpdateShipping(ctx context.Context, id uint, shippingId int) error
}
