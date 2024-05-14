package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

type OrderRepository interface {
	GetById(ctx context.Context, id int) (*model.Order, error)
	GetByIdAndStatus(ctx context.Context, id int, status model.OrderStatus) (*model.Order, error)
	Save(ctx context.Context, order model.Order) error
	UpdateStatus(ctx context.Context, id int, currentStatus, newStatus model.OrderStatus) error
	UpdateShipping(ctx context.Context, id int, shippingId int) error
}
