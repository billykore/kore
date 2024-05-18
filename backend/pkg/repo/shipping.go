package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

type ShippingRepository interface {
	GetById(ctx context.Context, id uint) (*model.Shipping, error)
	Save(ctx context.Context, shipping model.Shipping) (uint, error)
	UpdateStatus(ctx context.Context, id uint, newStatus, currentStatus model.ShippingStatus) error
}
