package order

import "context"

type Repository interface {
	GetById(ctx context.Context, id uint) (*Order, error)
	GetByIdAndStatus(ctx context.Context, id uint, status Status) (*Order, error)
	GetByShippingId(ctx context.Context, shippingId uint) (*Order, error)
	Save(ctx context.Context, order Order) error
	UpdateStatus(ctx context.Context, id uint, newStatus Status, currentStatus ...Status) error
	UpdateShipping(ctx context.Context, id uint, shippingId int) error
}
