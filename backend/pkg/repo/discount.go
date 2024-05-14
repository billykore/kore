package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

// DiscountRepository .
type DiscountRepository interface {
	List(ctx context.Context, limit int, startId int) ([]*model.Discount, error)
}
