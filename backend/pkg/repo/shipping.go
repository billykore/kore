package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

type ShippingRepository interface {
	Save(ctx context.Context, shipping model.Shipping) (uint, error)
}
