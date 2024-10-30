package shipping

import "context"

type Repository interface {
	GetById(ctx context.Context, id uint) (*Shipping, error)
	Save(ctx context.Context, shipping Shipping) (uint, error)
	UpdateStatus(ctx context.Context, id uint, newStatus, currentStatus Status) error
}
